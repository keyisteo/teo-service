
<h1>L I N I M A S A</h1>
{{range .Report}} 
	<div class="row justify-content-center">
		<div class="col-10" style="background-color: #f3f3f3	; padding: 15px;">
			<div class="row">
				<div class="col-4">
					<h3>Pelapor: {{(index $.Pelapor .IdReporter).Username}}</h3>
				</div>
				<div class="col-8 " align="right">
					<img src="/static/img/assets/kategori ({{.Category}}).png" style="max-height: 70px">
				</div>
			</div>
			<p>Tanggal Submit Laporan: {{ .CreatedAt}}</p>
			<img src="{{.LinkPhoto}}"	class="mx-auto d-block img-thumbnail" alt="Foto Laporan" style=" max-height: 400px;" />
			<h3>Detail Laporan: </h3>
			<p>{{ .Detail}}</p>
			<div class="row justify-content-center">
				<form action="/report/?id={{.Id}}" method="delete">
					<button class="btn btn-danger float-none" type="submit">Hapus</button>
				</form>
			</div>
		</div>
	</div>
{{end}}