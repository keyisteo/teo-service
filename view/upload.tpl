<html>
<head>
       <title>Upload file</title>
       <link rel="stylesheet" type="text/css" href="../css/image-picker.css">
       <script src="https://code.jquery.com/jquery-3.0.0.min.js" type="text/javascript"></script>
       <script type="text/javascript" src ="../js/image-picker.min.js"></script>
      	<link rel="stylesheet" type="text/css" href="../css/bootstrap.min.css">

</head>
<body>
	<div class="container">

		<form enctype="multipart/form-data" action="http://127.0.0.1:8080/upload" method="post">
			<input type="file" name="uploadfile" />
			<input type="text" name="name">
			<input type="text" name="deskripsi">m
			<select multiple="multiple" >
				<div class="row">
					<div class="col-4">
						<option data-img-src="../img/01.jpg" data-img-class="first" data-img-alt="Label 1" value="1" class="img-responsive">  Label 1  </option>
				  	</div>
				  	<div class="col-4">
				  		<option data-img-src="../img/02.jpg" data-img-alt="Label 2" value="2" class="img-responsive">  Label 2  </option>
				  	</div>
				  	<div class="col-4">
				  		<option data-img-src="../img/03.jpg" data-img-alt="Label 3" value="3" class="img-responsive">  Label 3 </option>
				  	</div>
				  	<div class="col-4">
				  		<option data-img-src="../img/04.jpg" data-img-alt="Label 4" value="4" class="img-responsive">  Label 4 </option>
					</div>
				  	<script type="text/javascript">
						// Initialize the object
					    $("select").imagepicker();
					    // Retrieve the picker
					    $("select").data('picker');
				    </script>
				</div>
			</select>
		    <input type="submit" value="upload" />
		</form>
	</div>
</body>
</html>