
	<h1>L I N I M A S A</h1>
	<!-- Gunakan range-->
		<div class="row justify-content-center">
			<div class="col-10" style="background-color: #f3f3f3	; padding: 15px;">
				<div class="row">
					<div class="col-4">
						<h3>Pelapor: {{ .Username}}</h3>
					</div>
					<div class="col-8 " align="right">
						<img src="{{/static/img/assets/kategori ({{ .Category}}).png" style="max-height: 70px">
					</div>
				</div>
				<p>Tanggal Submit Laporan: {{$report->created_at->diffForHumans()}}</p>
				<img src="/static/img/assets/{{ .LinkPhoto}}"	class="mx-auto d-block img-thumbnail" alt="Foto Laporan" style=" max-height: 400px;" />
				<div class="row">
					<div class="col-7">
						<h3>Detail Laporan: </h3>
						<p>{{$report->detail}}</p>
					</div>
				</div>

