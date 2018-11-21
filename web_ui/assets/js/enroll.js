var filter = /^([\w-\.]+)@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.)|(([\w-]+\.)+))([a-zA-Z]{2,4}|[0-9]{1,3})(\]?)$/;//

$(document).ready(function(){
     $("#start_test").click(function(event){
          if ($('#fname').val() == "") {
       	   alert('Please enter first name');
           event.preventDefault();
           return
       }

       if ($('#lname').val() == "") {
       	   alert('Please enter last name');
           event.preventDefault();
           return
       }

       if ($('#mobile').val() == "") {
       	   alert('Please enter mobile number');
           event.preventDefault();
           return
       }
 
        var sEmail = $('#email').val();
         if (!filter.test(sEmail)) {
            alert('Please enter valid email address');
            event.preventDefault();
            return
         }     

      $.post("enroll",
        JSON.stringify({ fist_name: $('#fname').val(), last_name: $('#lname').val(),
          mobile_no: $('#mobile').val(), email: $('#email').val() }),
		function(data, status){
	     if (status.toString() == "success" ) { 
	     var linkPage = document.getElementById('menu_item_starttest').href;
         window.location.href = linkPage + "#" + data;
        } else {
              alert("failed to start the test, please try again...")
           	  alert(data)
              event.preventDefault();
              return
		}
      });
  });
});