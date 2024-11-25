package page

import (
	"goapi/internal/api/repository/models"
)

func GetStaticData() []models.Person {
	people := []models.Person{
		{ID: 1, PersonID: "P001", TagID: "T001", PersonName: "Alice McDonals", RoomID: "100100"},
		{ID: 2, PersonID: "P002", TagID: "T002", PersonName: "Bob Hesburger", RoomID: "100100"},
		{ID: 3, PersonID: "P003", TagID: "T003", PersonName: "Jaakko Tuomisto", RoomID: "100100"},
		{ID: 4, PersonID: "P004", TagID: "T004", PersonName: "Jouko Aho", RoomID: "100100"},
		{ID: 5, PersonID: "P005", TagID: "T005", PersonName: "Teemu Tuomioja", RoomID: "100100"},
		{ID: 6, PersonID: "P006", TagID: "T006", PersonName: "Annikki Saari", RoomID: "100100"},
		{ID: 7, PersonID: "P007", TagID: "T007", PersonName: "Riikka Sirviö", RoomID: "100100"},
		{ID: 8, PersonID: "P008", TagID: "T008", PersonName: "Hannele Martinmäki", RoomID: "100100"},
		{ID: 9, PersonID: "P009", TagID: "T009", PersonName: "Ivan Svensson", RoomID: "100100"},
		{ID: 10, PersonID: "P010", TagID: "T010", PersonName: "Katja Suominen", RoomID: "100100"},
		{ID: 11, PersonID: "P011", TagID: "T011", PersonName: "Reima Eesti", RoomID: "100100"},
		{ID: 12, PersonID: "P012", TagID: "T012", PersonName: "Leo Suomi", RoomID: "100100"},
		{ID: 13, PersonID: "P013", TagID: "T013", PersonName: "Mona Lisa", RoomID: "100100"},
		{ID: 14, PersonID: "P014", TagID: "T014", PersonName: "Niina Veikkonen", RoomID: "100100"},
		{ID: 15, PersonID: "P015", TagID: "T015", PersonName: "Oskari Osaaja", RoomID: "100100"},
		{ID: 16, PersonID: "P016", TagID: "T016", PersonName: "Pauli Kunnas", RoomID: "100100"},
		{ID: 17, PersonID: "P017", TagID: "T017", PersonName: "Siiri Siikalatva", RoomID: "100100"},
		{ID: 18, PersonID: "P018", TagID: "T018", PersonName: "Veijo Japanilainen", RoomID: "100100"},
		{ID: 19, PersonID: "P019", TagID: "T019", PersonName: "Teemu Purra", RoomID: "100100"},
		{ID: 20, PersonID: "P020", TagID: "T020", PersonName: "Erkki Karvisto", RoomID: "100100"},
		{ID: 21, PersonID: "100", TagID: "69C3CBA3", PersonName: "Juha Tuomisto", RoomID: "-1"},
		{ID: 22, PersonID: "101", TagID: "47556663", PersonName: "Jenni Tuomarmäki", RoomID: "-1"},
		{ID: 23, PersonID: "102", TagID: "C4C28C4D", PersonName: "Saku Tenkula", RoomID: "-1"},
	}

	return people
}

func GetPageModel(data []*models.Person) string {
	persons := ""
	otherPersons := ""
	people := GetStaticData()
	for _, oPerson := range people {
		checked := false
		for _, person := range data {
			if oPerson.PersonID == person.PersonID {
				checked = true
				break
			}
			println(oPerson.PersonID + " ::: " + person.PersonID)
		}
		if !checked {
			otherPersons += get(oPerson.PersonName)
		}
	}
	for _, person := range data {
		persons += get(person.PersonName)
	}
	return `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Easy Access</title>
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
				<ul>` + otherPersons +
		`</ul>
			</div>
		</div>
	</body>
	</html>`
}

func get(data string) string {
	return "<li>" + data + "</li>"
}
