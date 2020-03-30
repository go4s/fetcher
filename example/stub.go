package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type ResponseBody struct {
	Code   int             `json:"code"`
	Reason string          `json:"reason"`
	Body   json.RawMessage `json:"body"`
}

func main() {
	g := gin.Default()
	g.GET("/json", func(ctx *gin.Context) {
		io.Copy(os.Stdout, ctx.Request.Body)
		defer ctx.Request.Body.Close()
		ctx.JSON(200, ResponseBody{
			Code:   200,
			Reason: "success",
			Body:   []byte(body),
		})
	})
	g.Run(":8080")
}

const body = `[
{
"_id": "5e6e42e8d278c7d660ff8546",
"index": 0,
"guid": "fd2a8b5f-6cd1-4f89-9a07-210f0d08ee2f",
"isActive": false,
"balance": "$3,343.64",
"picture": "http://placehold.it/32x32",
"age": 21,
"eyeColor": "brown",
"name": {
"first": "Curry",
"last": "Acosta"
},
"company": "GEEKY",
"email": "curry.acosta@geeky.biz",
"phone": "+1 (916) 579-2421",
"address": "361 Richards Street, Innsbrook, Ohio, 5892",
"about": "Laborum aliqua nisi aute labore enim anim veniam. Reprehenderit labore ullamco ex proident qui et aliqua ut. Id sit exercitation aliquip duis voluptate do proident officia excepteur. Elit et incididunt commodo amet. Consequat laborum id deserunt eu magna. Tempor veniam reprehenderit exercitation anim sunt minim eiusmod ad non fugiat elit ex commodo ea.",
"registered": "Friday, October 23, 2015 9:53 AM",
"latitude": "51.752467",
"longitude": "-85.62874",
"tags": [
"irure",
"irure",
"et",
"non",
"Lorem"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Salazar Woods"
},
{
"id": 1,
"name": "Madeline Barlow"
},
{
"id": 2,
"name": "Petersen Bridges"
}
],
"greeting": "Hello, Curry! You have 5 unread messages.",
"favoriteFruit": "strawberry"
},
{
"_id": "5e6e42e8d7b527c2e5fa2a4e",
"index": 1,
"guid": "1bd2fc41-00ae-4fb6-8d6d-cce6e828c52f",
"isActive": true,
"balance": "$3,120.02",
"picture": "http://placehold.it/32x32",
"age": 31,
"eyeColor": "brown",
"name": {
"first": "Orr",
"last": "Harrell"
},
"company": "DREAMIA",
"email": "orr.harrell@dreamia.net",
"phone": "+1 (820) 466-2168",
"address": "382 Kingston Avenue, Condon, North Carolina, 4215",
"about": "Voluptate esse exercitation id enim est. Ullamco sit exercitation laboris dolore dolor id quis labore quis Lorem non nostrud ex sit. Consectetur dolor magna occaecat non cupidatat sint aliqua id. Enim consectetur laborum exercitation aliqua exercitation in. Eu ut veniam eu elit incididunt ipsum do fugiat aliquip cillum cillum ipsum aliquip anim. Sit do cillum amet pariatur aliquip commodo occaecat reprehenderit ex in cillum id elit et.",
"registered": "Thursday, January 11, 2018 2:43 PM",
"latitude": "37.229135",
"longitude": "-16.514553",
"tags": [
"adipisicing",
"tempor",
"eiusmod",
"sit",
"nostrud"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Marion Hendrix"
},
{
"id": 1,
"name": "Haynes Gallegos"
},
{
"id": 2,
"name": "Hawkins Herrera"
}
],
"greeting": "Hello, Orr! You have 10 unread messages.",
"favoriteFruit": "strawberry"
},
{
"_id": "5e6e42e81b0aa9904bf8529d",
"index": 2,
"guid": "21174801-83ca-4e64-ad3d-c0dae8efd17a",
"isActive": false,
"balance": "$1,350.40",
"picture": "http://placehold.it/32x32",
"age": 21,
"eyeColor": "green",
"name": {
"first": "Tanisha",
"last": "Garner"
},
"company": "KOZGENE",
"email": "tanisha.garner@kozgene.name",
"phone": "+1 (836) 555-3461",
"address": "112 Montrose Avenue, Wintersburg, Vermont, 2309",
"about": "Ipsum sint qui duis sit dolor officia ex ut reprehenderit adipisicing. Cillum sint sunt in dolor mollit deserunt reprehenderit mollit laboris eu ad. Sint officia aute adipisicing non ipsum exercitation sint. Esse Lorem aliqua magna labore non fugiat cupidatat anim sint esse. Culpa Lorem voluptate excepteur et qui nisi laborum ullamco et.",
"registered": "Friday, April 1, 2016 7:42 PM",
"latitude": "17.787371",
"longitude": "-10.783008",
"tags": [
"nisi",
"anim",
"dolore",
"laboris",
"adipisicing"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Dunn Snider"
},
{
"id": 1,
"name": "Marguerite Mckee"
},
{
"id": 2,
"name": "James Valencia"
}
],
"greeting": "Hello, Tanisha! You have 6 unread messages.",
"favoriteFruit": "banana"
},
{
"_id": "5e6e42e868da8681212211f1",
"index": 3,
"guid": "b24387df-b71f-4f51-8518-feabc52f9f32",
"isActive": false,
"balance": "$1,367.99",
"picture": "http://placehold.it/32x32",
"age": 37,
"eyeColor": "green",
"name": {
"first": "Joanna",
"last": "Chapman"
},
"company": "SENMEI",
"email": "joanna.chapman@senmei.ca",
"phone": "+1 (926) 545-2191",
"address": "453 Farragut Place, Loma, North Dakota, 1755",
"about": "Reprehenderit nisi laborum esse consectetur nulla id velit ipsum reprehenderit laboris dolor anim Lorem. Elit velit deserunt ad enim aliquip consequat pariatur exercitation reprehenderit ea. Aliquip duis anim exercitation quis.",
"registered": "Wednesday, February 25, 2015 12:22 PM",
"latitude": "69.458074",
"longitude": "77.749258",
"tags": [
"est",
"tempor",
"laboris",
"aliqua",
"occaecat"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Williamson Maldonado"
},
{
"id": 1,
"name": "Shaffer Gilliam"
},
{
"id": 2,
"name": "Edwards Morrison"
}
],
"greeting": "Hello, Joanna! You have 9 unread messages.",
"favoriteFruit": "apple"
},
{
"_id": "5e6e42e870a2d03b0233ac8a",
"index": 4,
"guid": "01b41f87-b08c-4a59-8926-2cdfbc4eab1e",
"isActive": true,
"balance": "$3,631.18",
"picture": "http://placehold.it/32x32",
"age": 27,
"eyeColor": "brown",
"name": {
"first": "Maria",
"last": "Caldwell"
},
"company": "SPEEDBOLT",
"email": "maria.caldwell@speedbolt.tv",
"phone": "+1 (965) 429-3094",
"address": "499 Bushwick Court, Accoville, Puerto Rico, 2399",
"about": "Dolor adipisicing nostrud ullamco voluptate esse consequat qui. Fugiat labore laborum est culpa commodo tempor. Exercitation fugiat aliqua irure est excepteur ad nostrud veniam magna in ipsum. Irure adipisicing cupidatat duis non est adipisicing mollit id cupidatat ad. Irure nulla nulla in mollit sunt. Ad deserunt amet laboris do. Velit deserunt veniam sit ex irure.",
"registered": "Saturday, December 16, 2017 1:55 PM",
"latitude": "-32.495087",
"longitude": "-111.999844",
"tags": [
"Lorem",
"est",
"voluptate",
"exercitation",
"cupidatat"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Valarie Greer"
},
{
"id": 1,
"name": "Rogers Tyson"
},
{
"id": 2,
"name": "Elliott Hunter"
}
],
"greeting": "Hello, Maria! You have 6 unread messages.",
"favoriteFruit": "strawberry"
},
{
"_id": "5e6e42e82a0254dc28b9896d",
"index": 5,
"guid": "811bf19c-1dda-41df-bf5d-83d98bd269ce",
"isActive": false,
"balance": "$1,459.70",
"picture": "http://placehold.it/32x32",
"age": 31,
"eyeColor": "brown",
"name": {
"first": "Nancy",
"last": "Cabrera"
},
"company": "DYNO",
"email": "nancy.cabrera@dyno.com",
"phone": "+1 (951) 469-2753",
"address": "103 Dumont Avenue, Geyserville, Montana, 9181",
"about": "Elit deserunt culpa cupidatat magna anim quis culpa ex ipsum nulla sint. Pariatur est ex incididunt minim Lorem ea ad adipisicing qui. Id cupidatat sunt commodo pariatur culpa excepteur do et consequat anim. Irure amet anim cupidatat nulla nostrud non velit dolor consectetur incididunt pariatur. Consectetur reprehenderit adipisicing ad eu sunt enim.",
"registered": "Tuesday, June 3, 2014 3:20 PM",
"latitude": "-71.403009",
"longitude": "48.094775",
"tags": [
"anim",
"sint",
"ex",
"sunt",
"aliqua"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Catherine Richard"
},
{
"id": 1,
"name": "Bonnie Knight"
},
{
"id": 2,
"name": "Felecia Mcclure"
}
],
"greeting": "Hello, Nancy! You have 5 unread messages.",
"favoriteFruit": "apple"
},
{
"_id": "5e6e42e869fc0f289bb61bfa",
"index": 6,
"guid": "685e0339-425e-4c6e-a095-46565c2ecb89",
"isActive": false,
"balance": "$1,748.87",
"picture": "http://placehold.it/32x32",
"age": 21,
"eyeColor": "green",
"name": {
"first": "Knapp",
"last": "Dixon"
},
"company": "MANUFACT",
"email": "knapp.dixon@manufact.co.uk",
"phone": "+1 (927) 549-2851",
"address": "272 Williams Avenue, Gardners, Nevada, 3640",
"about": "Ut sint consequat reprehenderit est enim excepteur mollit elit fugiat cillum anim sit ea. Labore dolore exercitation nisi sunt ullamco Lorem do anim. Adipisicing nostrud tempor nostrud reprehenderit eiusmod est labore. Incididunt culpa voluptate amet consequat qui. Nostrud magna ut deserunt aliquip. Qui laboris culpa consectetur sunt sunt labore duis.",
"registered": "Thursday, May 31, 2018 2:47 AM",
"latitude": "2.362103",
"longitude": "-93.70473",
"tags": [
"culpa",
"nostrud",
"qui",
"exercitation",
"sunt"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Danielle Mercado"
},
{
"id": 1,
"name": "Kris Waters"
},
{
"id": 2,
"name": "Liz Collins"
}
],
"greeting": "Hello, Knapp! You have 10 unread messages.",
"favoriteFruit": "banana"
},
{
"_id": "5e6e42e8dd2b9c2d2abc6ae2",
"index": 7,
"guid": "f6c44060-25e6-4f66-b699-6b19118c4e9f",
"isActive": true,
"balance": "$1,045.59",
"picture": "http://placehold.it/32x32",
"age": 24,
"eyeColor": "blue",
"name": {
"first": "Barrera",
"last": "Miller"
},
"company": "EQUICOM",
"email": "barrera.miller@equicom.me",
"phone": "+1 (881) 600-3782",
"address": "909 Navy Street, Bartley, Wyoming, 7944",
"about": "Consequat aute officia pariatur ipsum anim occaecat tempor cillum occaecat adipisicing. Sunt consequat excepteur qui consequat consectetur fugiat cupidatat ullamco sint. Voluptate amet eu ex commodo elit. Nisi nisi est minim nostrud incididunt minim non cupidatat sit cillum nostrud qui. Enim veniam exercitation excepteur ullamco anim pariatur. Adipisicing ad adipisicing amet dolor mollit cupidatat.",
"registered": "Saturday, August 11, 2018 2:18 PM",
"latitude": "-71.010209",
"longitude": "-101.181986",
"tags": [
"culpa",
"ad",
"esse",
"aliqua",
"deserunt"
],
"range": [
0,
1,
2,
3,
4,
5,
6,
7,
8,
9
],
"friends": [
{
"id": 0,
"name": "Wolf Rivers"
},
{
"id": 1,
"name": "Landry Dillard"
},
{
"id": 2,
"name": "Jeanine Bowen"
}
],
"greeting": "Hello, Barrera! You have 10 unread messages.",
"favoriteFruit": "apple"
}
]`
