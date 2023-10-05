package main

import (
	"fmt"
	"strings"
)

func GetFoodView(food Food) string {
	var description string
	for _, desc := range strings.Split(food.Description, "\n") {
		description += "<li>" + desc + "</li>"
	}

	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1"/>
		<link rel="stylesheet" href="static/css/main.css">
		<title>Готовим %s</title>
	</head>
	<body>
	
	<div class="true_ans">
		<h1>Сегодня готовим:</h1>
		<h2>%s!</h2>
		<div class="div_image">
			<img src="%s" height="250px" max-width="400px" width="auto">
		</div>
		<h3><strong>Ингредиенты: </strong></h3>
		%s
		<li>
			<button onClick="window.location.reload();">Другое</button>
		</li>
	</div>
	</body>
	</html>`, food.Name, food.Name, food.Image, description)
}

func GetAddNewFood(food *Food) string {
	return fmt.Sprintf(`<!DOCTYPE html>
	<html lang="ru">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1"/>
		<title>GoRandFood</title>
	</head>
	<body>
	<div class="add_food">
		<h3>Добавление новой еды</h3>
		
		<p class="false_ans"><strong>Ошибка: </strong></p>
	
		<form action="" method="post" class="add_food_form" enctype="multipart/form-data">
			<p>
				<label for="password">Пароль для добавления еды</label>
				<input name="password" type="password">
			</p>
			<p>
				<label for="food_name">Название</label>
				<input name="food_name" type="text">
			</p>
			<p>
				<label for="food_descr">Описание</label>
				<input name="food_descr" type="text">
			</p>
			<p>
				<label for="food_image">Изображение</label>
				<input name="food_image" type="file" accept="image/png, image/jpeg">
			</p>
			<p>
				<input type="submit" value="Добавить">
			</p>
		</form>
	
	</div>
	
	</body>
	</html>`)
}
