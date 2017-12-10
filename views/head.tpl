<head>
    <title> {{ .Title}}</title>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <!-- Styles -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css" integrity="sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb" crossorigin="anonymous">

    <style type="text/css">
    	body { 
        padding-top: 70px; 
        margin-bottom: 60px;
      }
        
        .input-hidden {
          position: absolute;
          left: -9999px;
        }

        input[type=radio]:checked + label>img {
          border: 1px solid #fff;
          box-shadow: 0 0 3px 3px #090;
        }

        /* Stuff after this is only to make things more pretty */
        input[type=radio] + label>img {
          width: 150px;
          height: 150px;
          transition: 500ms all;
        }

        input[type=radio]:checked + label>img {
          transform: 
            rotateZ(-10deg) 
            rotateX(10deg);
        }

      html {
        position: relative;
        min-height: 100%;
      }

      .footer {
        position: absolute;
        bottom: 0;
        width: 100%;
        /* Set the fixed height of the footer here */
        height: 60px;
        line-height: 60px; /* Vertically center the text there */
        background-color: #f5f5f5;
      }

    </style>
</head>