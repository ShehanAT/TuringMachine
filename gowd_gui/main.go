package main

import (
	"github.com/dtylman/gowd"

	"fmt"
	"time"

	"github.com/dtylman/gowd/bootstrap"
)

var body *gowd.Element
var modal_div *gowd.Element
var close_modal_btn *gowd.Element

func main() {
	//creates a new bootstrap fluid container
	body = bootstrap.NewContainer(false)
	// add some elements using the object model
	div := bootstrap.NewElement("div", "well")
	row := bootstrap.NewRow(bootstrap.NewColumn(bootstrap.ColumnSmall, 3, div))
	body.AddElement(row)
	// add some other elements from HTML

	second_div := bootstrap.NewElement("div", "well")
	second_row := bootstrap.NewRow(bootstrap.NewColumn(bootstrap.ColumnSmall, 3, second_div))
	body.AddElement(second_row)

	// second_row.AddHTML(
	// 	`

	// 	<h1>Rules: </h1>
	// 	<div class="example1 wrapper">
	// 		<div class="main-item">width: 60%</div>
	// 		<div class="side-item nav">flex: 1</div>
	// 		<div class="side-item right-rail">flex: 1</div>
	// 	</div>
	// 	<!-- Button trigger modal -->

	// 	<!-- Modal -->
	// 	<div class="modal" id="modal-one">
	// 		<div class="modal-bg modal-exit"></div>
	// 		<div class="modal-container">
	// 			<h1>Amazing Modal</h1>
	// 			<h2>Pure Vanilla JavaScript</h2>
	// 			<button class="modal-close modal-exit">X</button>
	// 		</div>
	// 	</div>
	// 	`, nil)
	second_row.AddHTML(
		`
		<style>
			#create_rules_modal {
				position: fixed;
				background-color: rgba(255, 255, 255, 0.25);
				top: 0;
				right: 0;
				bottom: 0;
				left: 0;
				z-index: 999;
				
				opacity: 0;
				pointer-events: none;
				transition: all 0.3s;
				&:target {
				  visibility: visible;
				  opacity: 1;
				  pointer-events: auto;
				}
				& > div {
				  width: 400px;
				  position: absolute;
				  top: 50%;
				  left: 50%;
				  transform: translate(-50%, -50%);
				  padding: 2em;
				  background: white;
				}
				header {
				  font-weight: bold;
				}
				h1 {
				  font-size: 150%;
				  margin: 0 0 15px;
				}
			}
		</style>
		`, nil)

	// add a button to show a progress bar
	btn := bootstrap.NewButton(bootstrap.ButtonPrimary, "Start")
	close_modal_btn = gowd.NewElement("button")
	create_rule_btn := bootstrap.NewButton(bootstrap.ButtonPrimary, "Open Modal")
	btn.OnEvent(gowd.OnClick, btnClicked)
	create_rule_btn.OnEvent(gowd.OnClick, createRuleBtnClicked)
	close_modal_btn.OnEvent(gowd.OnClick, closeModalBtnClicked)

	modal_div = gowd.NewElement("div")

	heading_text := gowd.NewElement("h1")
	heading_text.SetText("Modal Text")
	close_modal_btn.SetText("Close")
	modal_div.AddElement(heading_text)
	modal_div.AddElement(close_modal_btn)
	modal_div.SetID("create_rules_modal")

	modal_div.Hide()
	row.AddElement(modal_div)
	row.AddElement(create_rule_btn)
	gowd.Run(body)
}

// happens when the 'start' button is clicked
func btnClicked(sender *gowd.Element, event *gowd.EventElement) {
	// adds a text and progress bar to the body
	sender.SetText("Working...")
	text := body.AddElement(gowd.NewStyledText("Working...", gowd.BoldText))
	progressBar := bootstrap.NewProgressBar()
	body.AddElement(progressBar.Element)

	// makes the body stop responding to user events
	body.Disable()

	// clean up - remove the added elements
	defer func() {
		sender.SetText("Start")
		body.RemoveElement(text)
		body.RemoveElement(progressBar.Element)
		body.Enable()
	}()

	// render the progress bar
	for i := 0; i <= 123; i++ {
		progressBar.SetValue(i, 123)
		text.SetText(fmt.Sprintf("Working %v", i))
		time.Sleep(time.Millisecond * 20)
		// this will cause the body to be refreshed
		body.Render()
	}

}

func createRuleBtnClicked(sender *gowd.Element, event *gowd.EventElement) {
	modal_div.Show()
	close_modal_btn.Show()
}

func closeModalBtnClicked(sender *gowd.Element, event *gowd.EventElement) {
	modal_div.Hide()
	close_modal_btn.Hide()
}
