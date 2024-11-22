package page

import (
	"goapi/internal/api/repository/models"
)

func GetPageModel(data []*models.Person) string {
	persons := ""
	for _, person := range data {
		persons += get(person.PersonName)
	}
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Kulunvalvontajärjestelmä MoveMe ACS</title>
		<style>
			body {
				width: auto;
				margin: 0;
				padding: 0;
				background-color: #F2EFE5;
				font-family: Arial, sans-serif; 
			}

			.box {
				background-color: #B4B4B8;
				padding: 20px;
				margin: 10px auto;
				/*border-radius: 8px;*/
				width: 100%;
				min-width: 200px;
				box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
			}

			.topBox {
				display: flex;
				gap: 10px;
				text-align: left; 
				padding: 5px 25px 5px 25px;
				box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1);
				background-color: #C7C8CC;
				width: auto;
			}

			.boxControl {
				display: flex;
				gap: 10px;
				width: calc(100% - max(40%, 20px));
				margin: 0 max(20%, 10px);
			}

			select, label {
				padding: 10px; 
				border-radius: 5px;
				border: 1px solid #ffe7a3;
				background-color: #fff;
				min-width: 200px;
				font-size: 1.0em;
				font-weight: bold; 
				color: #333;
				margin: 10px 0;
				text-align: center; 
				letter-spacing: 2px;
			}
			h1 {
				font-size: 1.5em;
				font-weight: bold; 
				color: #333;
				margin: 20px 0;
				text-align: center; 
				letter-spacing: 2px;
			}

			ul {
				list-style-type: none;
				padding: 0;
			}

			li {
				background-color: rgba(255, 255, 255, 0.6);
				padding: 10px;
				margin: 5px 0;
				border-radius: 5px;
				box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
			}
		</style>
	</head>
	<body>
		<div class="topBox">
			<select id="selection">
				<option value="option1">Opetustila 035b</option>
				<option value="option2">Opetustila 131</option>
				<option value="option3">Opetustila 331</option>
				<option value="option4">Opetustila 231</option>
			</select>
			<h1 style="opacity: 0.1">MoveMe ACS</h1>
		</div>
		<div class="boxControl">
			<div class="box" style="background-color: #E3E1D9;">
				<h1>Läsnäolijat</h1>
				<ul>` + persons +
		`</ul>
			</div>
			<div class="box" style="background-color: #E3E1D9;">
				<h1>Ei paikalla</h1>
				<ul>
				</ul>
			</div>
		</div>
	</body>
	</html>`
}

func get(data string) string {
	return "<li>" + data + "</li>"
}
