$(document).ready(function() {
 $.getJSON("admin", function(result) {
    var data =""
	for (var key in result) {
       if (result.hasOwnProperty(key)) {
           data = data.concat("<tr><td>", result[key].fist_name,"</td>");
		  data = data.concat("<td>", result[key].last_name,"</td>");
		  data = data.concat("<td>", result[key].email,"</td>");
		  data = data.concat("<td>", result[key].mobile_no,"</td>");
		  data = data.concat("<td>", result[key].taken_on,"</td>");
	      data = data.concat("<td>", result[key].score,"</td>");
	 	  data = data.concat('<td> <a href="/finishtest#',result[key].ref_no,'">',result[key].ref_no,"</a></td>"); 
	 }
}
   document.getElementById('mytable').innerHTML = data;	
 });


 $('#dtBasicExample').DataTable({
  "paging": true // false to disable pagination (or any other option)
 });
 $('.dataTables_length').addClass('bs-select');

});