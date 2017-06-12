/*
	{
		"created_on": "12 june 2017",
		"aim": "To develop REST APIs for iOS, Android and Web apps",
		"code_by": "Rishikesh Agrawani"
	}
*/

package main

import (
	"fmt"
	"net/http"
)


func Home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, htmlTextHome)
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

//HTML TEXTS
var htmlTextHome=`<!DOCTYPE html>
<html lang="en">
<head>
		<title>Profile</title>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
		<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
		<link href="https://maxcdn.bootstrapcdn.com/font-awesome/4.3.0/css/font-awesome.min.css" rel="stylesheet">
</head>
<body>
		<div class="container">
			<div class="row">
				<div class="col-md-12 text-center">
					<H1 style='color:green;' >Go is great</H1>
					<h5  style='color:blue;'>
						<span style='color:maroon;'>A post by</span> Rishikesh Agrawani
					</h5>
					<hr>
					<center>
						<img src='https://image.slidesharecdn.com/5bf6736d-861c-439d-b63b-0702cbb516de-151114034831-lva1-app6892/95/what-is-an-api-13-638.jpg?cb=1447944362' class="img-responsive" /><br>
						<img src='https://image.slidesharecdn.com/restapidesign-140729104120-phpapp02/95/rest-api-design-8-638.jpg?cb=1406675992' class="img-responsive" /><br>
						<img src="https://camo.githubusercontent.com/9b4e79cc86c2bfa29a11432d9d3ffb1ef6919545/687474703a2f2f6661726d392e737461746963666c69636b722e636f6d2f383139382f383233313837363431385f373264653064353737635f7a2e6a7067" class="img-responsive">
					</center>
				</div><!--Column-->
			</div><!--Row-->
		</div><!--Container-->
</body>
</html>`	
