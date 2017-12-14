<div class="container">
  <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
      <a class="navbar-brand" href="/">KESAT-ITB</a>
      <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarsExampleDefault" aria-controls="navbarsExampleDefault" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      
      <div class="collapse navbar-collapse" id="navbarsExampleDefault">
        <ul class="navbar-nav mr-auto">
        {{with .userData}}
          {{if eq .Type "2"}}
              <li class="nav-item">
                <a class="nav-link" href="/handle">Tanggapi</a>
              </li>
            <li class="nav-item">
              <a class="nav-link" href="/timeline">Linimasa</a>
            </li>
          {{else if eq .Type "1"}}
            <li class="nav-item">
              <a class="nav-link" href="/report">Laporan</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/timeline">Linimasa</a>
            </li>
          {{end}}
        {{end}}
        </ul>
        <ul class="nav navbar-nav navbar-right">
        {{with .userData}}
            <li class="nav-item dropdown"> 
              <a class="nav-link dropdown-toggle" href="/logout" id="dropdown01" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">{{.Username}}</a>
              <div class="dropdown-menu" aria-labelledby="dropdown01">
                <a class="dropdown-item" href="/logout">Logout</a>
              </div>
            </li>
            {{else}}
            <li class="nav-item">
              <a class="nav-link" href="/register">Register</a>
            </li>
            <li class="nav-item">
              <a class="nav-link" href="/login">Login</a>
            </li>
            {{end}}
        </ul>
      </div>
    </nav>
</div>
