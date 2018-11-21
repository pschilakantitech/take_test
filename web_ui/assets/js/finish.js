$( document ).ready(function() {
	$("#mitem_test_started").hide();
    var refNo = window.location.hash.substring(1)
	$.getJSON("getresult?ref_no="+refNo , function(data){
        $("#fist_name").html(data.fist_name);
		$("#last_name").html(data.last_name);
		$("#mobile_no").html(data.mobile_no);
		$("#email").html(data.email);
		$("#score").html(data.score);
		$("#taken_on").html(data.taken_on);
		$("#started_at").html(data.started_at);
		$("#ended_at").html(data.ended_at);
		$("#duration").html(data.duration);
		$("#ref_no").html(refNo);
		
   });
});
  