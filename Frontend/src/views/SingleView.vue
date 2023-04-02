<template>  
   <div v-if="isLoggedIn" class="single">        
    <form id="form" style="display:flex;flex-direction:row;align-items:center;justify-content:center;">
      <label for="eq">Enter an equation:</label>
      <input style="margin:0px 20px;" type="text" id="eq" v-model="expression">
      <input type="button" id="submit-button" value="Draw" @click="trigger($event)">
    </form>

    <div id="plot"></div>

    <div id="visualizer-audio">
      <section id="section1">
        <div style="display:flex;flex-direction:column;align-items:center;justify-content:space-between;">
          <midi-player style="margin-top:20px;" :src="midiurl" ref="src3">
          </midi-player>
          <midi-visualizer style="margin-top:20px;" ref="src1" :src="midiurl" type="staff">        
          </midi-visualizer>
          <midi-visualizer ref="src2" :src="midiurl" visualizer="#section1 midi-visualizer">
          </midi-visualizer>
        </div>
      </section>
    </div>
  </div>
</template>

<script lang="ts">
import { mapGetters } from "vuex";
export default {
  data() {
    return {
      midiurl: "",
      expression: ""
    }
  },  
  computed: {
    ...mapGetters(["isLoggedIn", "token"])
  },
  mounted() {
      const visualizerSettings = {
        noteHeight: 3,
        pixelsPerTimeStep: 60,
        minPitch: 21
      };

      const visualizerSettings2 = {
        noteHeight: 6,
        pixelsPerTimeStep: 90,
        minPitch: 21
      };

      this.$refs.src1.config = visualizerSettings
      this.$refs.src2.config = visualizerSettings2
      this.$refs.src3.config = visualizerSettings

      let script1 = document.createElement('script')
      let script2 = document.createElement('script')
      let script3 = document.createElement('script')
      script1.setAttribute('src', 'https://unpkg.com/mathjs@11.7.0/lib/browser/math.js')
      script2.setAttribute('src', 'https://cdn.plot.ly/plotly-1.35.2.min.js')
      script3.setAttribute('src', 'https://cdn.jsdelivr.net/combine/npm/tone@14.7.58,npm/@magenta/music@1.23.1/es6/core.js,npm/focus-visible@5,npm/html-midi-player@1.5.0')
      document.head.appendChild(script1)
      document.head.appendChild(script2)
      document.head.appendChild(script3)
  },
  methods: {
    trigger() {
      const clamp = (num:any, min:any, max:any) => Math.min(Math.max(num, min), max);
      const min = 21;
      const max = 108;
      const self = this;
      try {
        // Clear existing midi
        this.midiurl = "";

        // compile the expression once
        const expression = this.expression;
        
        const expr = math.compile(expression)

        // evaluate the expression repeatedly for different values of x
        
        const xValues = math.range(-10, 11, 1).toArray()
        const yValues = xValues.map(function (x:any) {
          //return expr.evaluate({x: x})
          return clamp(Math.round(expr.evaluate({x: x})), min, max)
        })

        // render the plot using plotly
        const trace1 = {
          x: xValues,
          y: yValues,
          type: 'scatter'
        }
        const data = [trace1]
        
        Plotly.newPlot('plot', data)

        console.log(this.expression)
        fetch(import.meta.env.VITE_API_URL+"midi",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json", 
            "Authorization": "Bearer " + self.token         
          },        
          body: JSON.stringify({
            Fn: btoa(self.expression),
            NotesTime: xValues,
            NotesValue: yValues
          }),
        }).then(function() {
          self.midiurl = "http://localhost:8082/"+btoa(self.expression)+".mid";          
        }).catch(function(err) {
          console.log('Fetch Error: ', err);
        });
      }
      catch (err) {
        console.error(err)
        alert(err)
      }

    }
  }
}
</script>

<style>
input[type=text] {
      width: 300px;
}
input {
  padding: 6px;
}
body, html, input {
  font-family: sans-serif;
  font-size: 11pt;

}
form {
  margin: 20px 0;
}
</style>