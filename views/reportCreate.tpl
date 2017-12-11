
	

	<script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.1/jquery.min.js"></script>
	<h1>L A P O R K A N</h1>
	<div class="row justify-content-center">
		<div class="col-10">
			<form runat="server" class="form-horizontal" method="POST" action="/report" enctype="multipart/form-data" id="form_report">
		    	<input required type="hidden" class="form-control" name="id_reporter">
				<div class="form-group">
				   	<label for="picture"><h5>Foto Laporan:</h5></label>
					<input required onchange="document.getElementById('blah').src = window.URL.createObjectURL(this.files[0])" type="file" id="imgInp" class="form-control" name="picture"  >
					<img id="blah" class="card img-thumbnail rounded mx-auto d-block" alt="Foto Laporan" src="/static/img/assets/image-placeholder.png" />
				</div>
				<div class="form-group">
					<label for="detail"><h5>Deskripsi:</h5></label>
					<textarea required name="detail" form="form_report" class="form-control" rows="5" placeholder="Tuliskan Keluhan Anda Di Sini"></textarea>
				</div>
				<div class="form-group">
					<label for="category"><h5>Kategori:</h5></label>
					<br>
					<div class="row justify-content-center">
						<div>
						    <input class="input-hidden" hidden type="radio" id="category1" name="category" value="1">
						    <label for="category1">
						    	<img src="/static/img/assets/kategori (1).png" alt="Kebersihan" />
						    </label>
						</div>
						<div>
						    <input class="input-hidden" type="radio" id="category2" name="category" value="2">
						    <label for="category2">
						    	<img src="/static/img/assets/kategori (2).png" alt="Kelas" />
						    </label>
						</div>
						<div>
						    <input class="input-hidden" type="radio" id="category3" name="category" value="3">
						    <label for="category3">
						    	<img src="/static/img/assets/kategori (3).png" alt="Publikasi" />
						    </label>
						</div>
					</div>
					<div class="row justify-content-center">
						<div>
						    <input class="input-hidden" type="radio" id="category4" name="category" value="4">
						    <label for="category4">
						    	<img src="/static/img/assets/kategori (4).png" alt="Ruang dan Gedung" />
						    </label>
						</div>
						<div>
						    <input class="input-hidden" type="radio" id="category5" name="category" value="5">
						    <label for="category5">
						    	<img src="/static/img/assets/kategori (5).png" alt="Transportasi" />
						    </label>
						</div>
						<div>
						    <input class="input-hidden" type="radio" id="category6" name="category" value="6">
						    <label for="category6">
						    	<img src="/static/img/assets/kategori (6).png" alt="Umum" />
						    </label>
					    </div>
				    </div>
				</div>
				<button class="btn btn-primary" type="submit">Kirim</button>
			</form>
		</div>
	</div>

	<script type="text/javascript">
		function readURL(input) {
		  if (input.files && input.files[0]) {
		    var reader = new FileReader();

		    reader.onload = function(e) {
		      $('#picture-show').attr('src', e.target.result);
		      $('#picture-show').attr('style', "");
		    }

		    reader.readAsDataURL(input.files[0]);
		  }
		}

		$("#imgInp").change(function() {
		  readURL(this);
		});
	?>
	</script>
