<?php
// // setcookie(name, value, expire, path, domain, secure, httponly);
// setcookie("Role","Ad4miN", time() +1,"/" , "localhost",  "secure" , "httponly");

// print_r($_COOKIE);
function test_input($data) {
    $data = trim($data);
    $data = stripslashes($data);
    $data = htmlspecialchars($data);
    return $data;
  } 

 



?>

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Sign in</title>

    <!-- <form action="" method="get">

    <label for="mail"></label>
<input type="text" name="mail" id="mail ">
<button>submit</button> -->


    <!-- </form> -->
<center>
<!-- <form action="" method="post"  >





<label for="username">Enter Your User Name : </label>
<input type="text" name="username" id="username" required>
<br><br>

<label for="password">Enter Your Password : </label>
<input type="password" name="password" id="password" required>
<br><br>
<button id="loginbtn" name="loginbtn" >Login</button> -->







</form>
</center>






</head>
<body>












</body>
</html>