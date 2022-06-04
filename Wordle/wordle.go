package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	title  string = "Wordle"
	width  int    = 435
	height int    = 600
	rows   int    = 6
	cols   int    = 5
)

type ECheckType int

const (
	ECT_None ECheckType = iota
	ECT_Correct
	ECT_HalfCorrect
	ECT_Wrong
)

var (
	fontsize    = 32
	dpi         = 72
	normal_font font.Face

	background_color = color.White
	lightgray        = color.RGBA{0xc2, 0xc5, 0xc6, 0xff}
	grey             = color.RGBA{0x77, 0x7c, 0x7e, 0xff}
	yellow           = color.RGBA{0xcd, 0xb3, 0x5d, 0xff}
	green            = color.RGBA{0x60, 0xa6, 0x65, 0xff}
	font_color       = color.Black

	edge     bool   = false
	alphabet string = "abcdefghijklmnopqrstuvwxyz"
	grid     [cols * rows]string
	dict     []string
	check    [cols * rows]int
	location int  = 0
	is_won   bool = false
	answer   string
)

type Game struct {
	runes []rune
}

func KeyPressed(Key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)

	duration := inpututil.KeyPressDuration(Key)
	if duration == 1 {
		return true
	}

	if duration >= delay && (duration-delay)&interval == 0 {
		return true
	}

	return false
}

func (game *Game) Update() error {
	if !is_won {
		game.runes = ebiten.AppendInputChars(game.runes[:0])

		if strings.Contains(alphabet, string(game.runes)) && string(game.runes) != "" && location >= 0 && location < rows*cols {
			grid[location] = string(game.runes)[0:1]
			if !edge {
				location++
			}
		}
	}

	edge = false

	if (location+2)%cols == 1 && location != 0 {
		edge = true
	}

	if edge == true && KeyPressed(ebiten.KeyEnter) {
		input := ""
		for i := (location - (cols - 1)); i < (location + 1); i++ {
			input += grid[i]
		}
		// fmt.Print(input)

		is_valid_word := false
		for _, word := range dict {
			if word == input {
				is_valid_word = true
			}
		}

		if is_valid_word {
			var check_word [cols]bool
			for i, letter := range input {
				for j, ans_letter := range answer {
					if i == j && string(letter) == string(ans_letter) {
						check[location-cols+i+1] = int(ECT_Correct)
					} else {
						check[location-cols+i+1] = int(ECT_Wrong)
					}
				}
			}

			for i, letter := range input {
				if strings.Contains(answer, string(letter)) {
					found := false
					for j, ans_letter := range answer {
						if found == false && check_word[j] == false && check[location-cols+i+1] == 0 {
							if string(letter) == string(ans_letter) {
								check_word[j] = true
								found = true
								check[location-cols+i+1] = int(ECT_HalfCorrect)
							}
						}
					}
				}
			}

			if input == answer {
				is_won = true
			}
			location++
			edge = false
		}
	}

	if KeyPressed(ebiten.KeyBackspace) && location > 0 {
		if check[location-1] == 0 {
			grid[location] = ""
			location--
		}
	}

	if location < 0 {
		location = 0
	}

	if location > rows*cols {
		location = rows*cols - 1
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(background_color)

	if is_won {
		winner := "Good Job!!"
		for i := 0; i < len(winner); i++ {
			msg := fmt.Sprintf(strings.ToUpper(string([]rune(winner)[i])))
			font_color = color.Black

			text.Draw(screen, msg, normal_font, i*85+40, rows*85+55, font_color)
		}
	}

	for w := 0; w < cols; w++ {
		for h := 0; h < rows; h++ {
			var loc = w + h*cols

			rect := ebiten.NewImage(75, 75)
			rect.Fill(lightgray)
			font_color = color.Black

			if check[loc] != int(ECT_None) {
				switch check[loc] {
				case int(ECT_Correct):
					rect.Fill(green)
				case int(ECT_HalfCorrect):
					rect.Fill(yellow)
				case int(ECT_Wrong):
					rect.Fill(grey)
				}
				font_color = color.White
			}

			if loc == location && check[loc] == int(ECT_None) {
				rect.Fill(grey)
			}

			option := &ebiten.DrawImageOptions{}
			option.GeoM.Translate(float64(w*85+10), float64(h*85+10))
			screen.DrawImage(rect, option)

			if check[loc] == int(ECT_None) {
				rect2 := ebiten.NewImage(73, 73)
				rect2.Fill(color.White)

				option2 := &ebiten.DrawImageOptions{}
				option2.GeoM.Translate(float64(w*85+10)+1, float64(h*85+10)+1)
				screen.DrawImage(rect2, option2)
			}

			if grid[loc] != "" {
				msg := fmt.Sprintf(strings.ToUpper(grid[loc]))
				text.Draw(screen, msg, normal_font, w*85+38, h*85+55, font_color)
			}
		}
	}

	if !is_won && check[len(check)-1] != int(ECT_None) {
		for i := 0; i < len(answer); i++ {
			msg := fmt.Sprintf(strings.ToUpper(string([]rune(answer)[i])))
			font_color = color.Black

			text.Draw(screen, msg, normal_font, i*85+40, rows*85+55, font_color)
		}
	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	game := &Game{}
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		log.Fatal(err)
	}

	normal_font, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontsize),
		DPI:     float64(dpi),
		Hinting: font.HintingFull,
	})

	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle(title)

	content, err := ioutil.ReadFile("dict.txt")
	if err != nil {
		log.Fatal(err)
	} else {
		dict = strings.Split(string(content), "\n")
	}

	rand.Seed(time.Now().UnixNano())
	answer = dict[rand.Intn(len(dict))]
	fmt.Printf(answer)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
