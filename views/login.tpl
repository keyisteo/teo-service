
	<h1>L O G I N</h1>
	<br>
	<br>
	<div class="row justify-content-center">
        <div class="col-md-offset-5 col-md-3">
            <form action="/login" method="POST" class="form-horizontal">
            <div class="form-login">
            {{ .xsrfdata }}
            <h4>Welcome back.</h4>
            <input name="username" type="text" id="userName" class="form-control input-sm chat-input" placeholder="Username" />
            </br>
            <input name="password" type="password" id="userPassword" class="form-control input-sm chat-input" placeholder="Password" />
            </br>
			<button class="btn-primary" type="submit">Login</button>
            </div>
        	</form>
        </div>
    </div>

