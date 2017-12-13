
	<h1 align="center">R E G I S T E R</h1>
		<br>
		<br>
	<div class="row justify-content-center">

    	<div class="col-md-offset-5 col-md-3">
			<form action="/register" method="POST" >
					 {{ .xsrfdata }}
					<div class="form-group">
						<label class="form-label" for="username"> Username </label>
						<input class="form-control" required type="text" name="username" id="username">
					</div>
					<div class="form-group">
						<label class="form-label" for="email"> Email </label>
						<input class="form-control" required type="email" name="email" id="email">
					</div>
					<div class="form-group">
						<label class="form-label" for="password"> Password </label>
						<input class="form-control" required type="password" name="password" id="password">
					</div>
					<div class="form-group">
						<label class="form-label" for="password_confirmation"> Password Confirmation</label>
						<input class="form-control" required type="password" name="password_confirmation" id="password_confirmation">
					</div>
					<button class="btn-primary" type="submit">Register</button>
				</div>
			</form>
		</div>
	</div>

