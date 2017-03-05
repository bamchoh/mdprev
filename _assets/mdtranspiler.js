$(function(){
	var test_input = $("#text-input")
	function DoPreview() {
		marked.setOptions({
			langPrefix: 'hljs ',
			highlight: function (code) {
				return hljs.highlightAuto(code).value;
			}
		});
		preview.innerHTML = marked(test_input.val());
	}

	$('#text-input').keyup(function() {
		DoPreview();
	});

	$('#text-input').change(function(){
		DoPreview();
	});

	function InitTextInput() {
		$.ajax({
			type:"get",
			url: "/api/markdown",
			contentType: 'application/json',
			dataType: "json",
			success: function(json_data) {
				var data = JSON.parse(json_data);
				test_input.val(data["md"]);
			},
			error: function() {
				test_input.val("Write **Markdown** in the field.");
			},
			complete: function() {
				DoPreview();
				var data = test_input.val();
				test_input.attr('data-mk', data);
			}
		});
	}
	InitTextInput();

	setInterval(function() {
		var prev = test_input.attr('data-mk');
		var md_val = test_input.val();
		if (prev == md_val) {
			return
		}
		test_input.attr('data-mk',md_val);

		var data = {
			md: md_val
		};

		$.ajax({
			type:"post",
			url: "/api/markdown",
			data: JSON.stringify(data),
			contentType: 'application/json',
			dataType: "json",
			success: function(json_data) {
				console.log(json_data)
			},
			error: function() {
			},
			complete: function() {
			}
		});
	},5000);
});
