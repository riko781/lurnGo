<?php

$timestart=microtime(true);

$con = mysqli_connect("192.168.1.147","inovatic","sysiphe","production");

// Check connection
if (mysqli_connect_errno())
  {
  echo "Failed to connect to MySQL: " . mysqli_connect_error();
  }
for ($i=0;$i<1000000;$i++){
    $request="SELECT rame FROM info_bilan WHERE id=5375495";

    $result=mysqli_query($con,$request);

    $row=mysqli_fetch_assoc($result);
   
}

$timeend=microtime(true);
$time=$timeend-$timestart;
$page_load_time = number_format($time, 3);
echo "<br>Script execute en " . $page_load_time . " sec";
