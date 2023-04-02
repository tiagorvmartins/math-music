<template>
    <div class="centered login-wrapper">
        <h1>LOGIN</h1>
        <form @submit.prevent="login">
            <input v-model="username" placeholder="username" />
            <br />
            <br />
            <input v-model="password" placeholder="password" type="password" />
            <br />
            <br />
            <div class="centered">
                <button type="submit">Login</button>
            </div>
        </form>
    </div>
</template>


<script lang="ts">
import { mapMutations } from "vuex";
export default {
    data: () => {
        return {
            username: "",
            password: "",
        };
    },
    methods: {
        ...mapMutations(["setUser", "setToken"]),
        async login(e : any) {
            e.preventDefault();
            const response = await fetch(import.meta.env.VITE_API_URL+"login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    User: this.username,
                    Password: this.password,
                }),
            });
            const { user, token } = await response.json();
            this.setUser(user);
            this.setToken(token);
            this.$router.push("/link");
        },
    }    
};
</script>

<style scoped>
.login-wrapper {
    flex-direction: column;
}

.centered {
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>