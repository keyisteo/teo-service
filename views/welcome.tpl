
{{with .userData}}

<br>    
<br>    
<div  class="row justify-content-center">
    <div class="col-8">
        <h1 align="center">S E L A M A T</h1>
    </div>
</div>    
<br>      
<div class="row justify-content-center">
    <div class="col-8">
        <h1 align="center">D A T A N G</h1>
    </div>
</div>
<br>  
<div class="row justify-content-center">
    <div class="col-8">
        <h1 align="center">{{.Username}}</h1>
    </div>
</div>  
<br>    
<br>    
<div class="row justify-content-center">
    <div class="col-8">
        <h5 align="center">K E S A T  I T B</h5>
    </div>
</div>
<br>   


{{else}}

<br>    
<br>    
<div  class="row justify-content-center">
    <div class="col-8">
        <h1 align="center">S E L A M A T</h1>
    </div>
</div>    
<br>      
<div class="row justify-content-center">
    <div class="col-8">
        <h1 align="center">D A T A N G</h1>
    </div>
</div>
<br>    
<br>    
<br>    
<div class="row justify-content-center">
    <div class="col-8">
        <h5 align="center">K E S A T  I T B</h5>
    </div>
</div>
<br>    
<div class="row justify-content-center">
    <div align="center" class="col-8">
        <form action="/login" method="get">
            <button class="btn btn-secondary float-none" type="submit" align>Login</button>
        </form>
    </div>
</div>
<br>       
</div>
<div align="center" class="row justify-content-center">
    <div class="col-8">
        <form action="/register" method="get">
            <button class="btn btn-primary float-none" type="submit" align>Register</button>
        </form>
    </div>
</div>
{{end}}