var totalQuestions = 0;

$( document ).ready(function() {
  var refNo = window.location.hash.substring(1)
  $.getJSON("starttest?ref_no="+refNo, function(data) {
          $("#question").html( data.serial_no+") " +data.question);  
          $("#optionA").html(data.option_a);
          $("#optionB").html(data.option_b);
          $("#optionC").html(data.option_c);
          $("#optionD").html(data.option_d);
          $("#ref_no").html(data.ref_no);
          $("#question_id").html(data.question_id); 
		  $("#serial_no").html(data.serial_no);
    }); 
    $.getJSON("testdetails" , function(result){
  		totalQuestions=result.test_questions;
     });	

    $("#nextQuestion").click(function(event){
 	var option = $("input[type='radio'][name='option']:checked").val();
        if (option === undefined){
       		alert("please select an answer.")
	    event.preventDefault();
	    return 
	 }
 
	var slNo=$('#serial_no').text(); 
    var updateURL="updateans?ref_no=" + $('#ref_no').text() + "&q_id=" + $('#question_id').text() + "&ans=" + option;
    $.getJSON(updateURL , function(data) {
          alert("got next q")  
         $("#question").html( data.serial_no+") " +data.question);  
           $("#optionA").html(data.option_a);
           $("#optionB").html(data.option_b);
           $("#optionC").html(data.option_c);
           $("#optionD").html(data.option_d);
           $("#ref_no").html(data.ref_no);
           $("#question_id").html(data.question_id);
      	   $("#serial_no").html(data.serial_no);
	 });
   	if ( slNo == totalQuestions) {
         var linkPage = document.getElementById('mitem_finshtest').href;
         window.location.href = linkPage + "#" + refNo;
    } 
  
  });	
});