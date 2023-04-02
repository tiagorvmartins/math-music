import json
import sys
from midiutil.MidiFile import MIDIFile

with open(sys.argv[1], 'r') as json_file:
    json_data = json.load(json_file)    

    # create your MIDI object
    mf = MIDIFile(1)     # only 1 track
    track = 0   # the only track

    time = 0    # start at the beginning
    mf.addTrackName(track, time, "Sample Track")
    mf.addTempo(track, time, 120)

    # add some notes
    channel = 0
    volume = 100

    index = 0
    time = 0
    duration = 1
    for y in json_data["NotesValue"]:
       pitch = y
       mf.addNote(track, channel, pitch, time, duration, volume)   
       time = time + duration

    # write it to disk
    with open("static/"+sys.argv[2]+".mid", 'wb') as outf:
        mf.writeFile(outf)
        print("Wrote MIDI file with success!")