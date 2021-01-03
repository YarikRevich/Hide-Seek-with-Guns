package Sound

import (
	"os"
	"time"
	"Game/Components/States"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Sound interface{
	//Interface for sound playing.

	Init(*States.States)
	Play()
}

type S struct{
	//Structure for playing sound which saves states

	states *States.States
}

func (s *S) Init(states *States.States){
	//Inits states for structure

	s.states = states
}

func (s *S) Play(){
	//Plays places soundtrack

	if !s.states.ComponentsStates.PlayGameSound{
		s.states.ComponentsStates.PlayGameSound = true
		go func(){
			sound, err := os.Open("StartSound.mp3")
			if err != nil{
				panic(err)
			}
			defer sound.Close()
			streamer, format, err := mp3.Decode(sound)
			if err != nil{
				panic(err)
			}
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			done := make(chan bool)
			speaker.Play(beep.Seq(streamer, beep.Callback(func(){
				done <- true
			})))
			<- done
			s.states.ComponentsStates.PlayGameSound = false
		}()
	}
}