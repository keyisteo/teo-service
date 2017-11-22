<!DOCTYPE html>
<html>
<head>
	<title>Firebase Web Basics</title>

	<link href="https://fonts.googleapis.com/css?family=Raleway:300,400,500,700" rel="stylesheet">

	<script src="https://use.fontawesome.com/939e9dd52c.js"></script>

	<link rel="stylesheet" type="text/css" href="style.css">
</head>
<body>

	<div class="mainDiv" align="right">
		<h1 align="left">Firebase Web App</h1>
		<textarea placeholder="Enter text here..."></textarea>
		<button><i class="fa fa-arrow-right" aria-hidden="true"></i></button>		
	</div>


	<script src="https://www.gstatic.com/firebasejs/4.6.2/firebase.js"></script>
	<script>
	  // Initialize Firebase
	  var config = {
	    apiKey: "AIzaSyBM6otKpgUnqhDWvhAUw4Hgzb7b-VEKKfs",
	    authDomain: "teo-service.firebaseapp.com",
	    databaseURL: "https://teo-service.firebaseio.com",
	    projectId: "teo-service",
	    storageBucket: "teo-service.appspot.com",
	    messagingSenderId: "108178798942"
	  };
	  firebase.initializeApp(config);
	</script>
	<script type="text/javascript">window.alert("okay");</script>
</body>
</html>

