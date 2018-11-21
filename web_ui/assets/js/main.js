 
$( document ).ready(function() {
	$("#mitem_test_started").hide();
    $.getJSON("testdetails" , function(result){
        $("#total_qustions").text("Total number of questions: "+ result.test_questions);
        $("#total_time").text("Time alloted: "+result.test_time);
       });
});
 
