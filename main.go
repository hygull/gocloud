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

func Chart(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, htmlTextChart)
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/charts/", Chart)

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

var htmlTextChart = `
<!DOCTYPE html>
<html>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">

  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css">
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.1/jquery.min.js"></script>
  <script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  <script src="https://ajax.googleapis.com/ajax/libs/angularjs/1.4.8/angular.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.1.4/Chart.min.js"></script>
<!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.2.1/Chart.bundle.js"></script> -->
  <style type="text/css">
    h2{
      color: lightgreen;
    }
    p{
      color: blue;
    }
    h2,p{
      text-align: center;
    }
    h1{
      color:white;
    }
    th{
      font-size: 20px;
      background-color: #64453F;
    }
    td{
      font-size: 15px;
    }
    td:nth-child(even){
      background-color: gray;
            color:white;
    }
    td:nth-child(odd){
      background-color: pink;

    }
    th{
      font-weight: bolder;
      color:white;
    }
    td{
      font-weight: bold; 
    }
    body{
     /*background-color:#371E1E;*/
    }
    p{
      color: pink;
    }
    #a_link{
      color: lightblue;
    }
  </style>
<body>

<div class="container">
<!-- <center><a role="button" class="btn btn-info" href = "/api/student/mathematics/" target="_blank">Add result for a student </a></center><br> -->
    <center><h1 style="color: green">Line Chart</h1></center>
    <hr>
    <br>
        <div class="row">
            <div class="col-md-6">
              <canvas id="myChart1"></canvas>
              <br>
              <hr>
            </div>
            <div class="col-md-6">
              <canvas id="myChart2"></canvas>
              <br>
              <hr>
            </div>
              
        </div>
        <div class="row">
            <div class="col-md-6">
              <canvas id="myChart3"></canvas>
              <br>
              <hr>
            </div>
            <div class="col-md-6">
              <canvas id="myChart4"></canvas>
              <br>
              <hr>
            </div>
              
        </div>

</div>

<!---->

<!---->

<br>
<!-- <center><h4><a role="button" class="btn btn-success" href = "/api/student/mathematics/" target="_blank" class="btn btn success" >Add results for a student </a></h4></center><br> -->

<script type="text/javascript">
  
  // var jsonData = $.ajax({
  //   // url: 'http://filtron.pythonanywhere.com/api/auth-users/',
  //   // url:'http://127.0.0.1:8000/api/auth-users/',
  //   //url:"http://127.0.0.1:8000/api/student/mathematics/",
  //   url:"http://filtron.pythonanywhere.com/api/student/mathematics/",
  //   dataType: 'json',
  // }).done(function (results) {
      var typeArr = ["pie", "bar", "line", "polarArea"]

      for(var i=1;i<=4;i++){

            var ctx = document.getElementById("myChart"+i).getContext("2d");
            var myChart = new Chart(ctx,{
                type: typeArr[i-1],
                data:{
                  // labels:firstnames,
                  labels:["enquiries","orders","deleveries","invoices","receipts"],
                  
                  datasets:[
                        {
                          label:"",
                          data:[56, 67, 89, 90, 57],
                          backgroundColor:[ "darkgray", "red", "green", "blue", "pink"]
                        },

                    ],
                  options: {
                      title: {
                          display: true,
                          text: 'My Users Chart',
                          fontColor:"navy"
                      }
                  }
                }
            });
        }
          //Char js code ends

         
  // })
//-->
</script>
</script>
</body>
`
