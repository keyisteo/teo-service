
<h1>L A P O R A N</h1>
	<div class="row justify-content-center">
		<div class="col-10" style="background-color: #f3f3f3	; padding: 15px;">
			<div class="row">
				<div class="col-4">
					<h3>Pelapor: {{ .Report.IdReporter}}</h3>
				</div>
				<div class="col-8 " align="right">
					<img src="/static/img/assets/kategori ({{.Report.Category}}).png" style="max-height: 70px">
				</div>
			</div>
			<p>Tanggal Submit Laporan: {{ .Report.CreatedAt}}</p>
			<img src="{{.Report.LinkPhoto}}"	class="mx-auto d-block img-thumbnail" alt="Foto Laporan" style=" max-height: 400px;" />
			<h3>Detail Laporan: </h3>
			<p>{{ .Report.Detail}}</p>
		</div>
	</div>