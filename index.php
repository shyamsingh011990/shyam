<html>

<head>
<title>titpush using node</title>
<script src="https://www.gstatic.com/firebasejs/5.4.1/firebase.js"></script>
<link rel="manifest" href="manifest.json">


<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.3.1/jquery.min.js"></script>



<!-- <script src="./client.js"></script>
 -->
</head>
<body>
<script>
  var config = {
    apiKey: "AIzaSyDSbWYLvRn1SjDGyieNJgWJcFaLdkSAmvo",
    authDomain: "mani-83d3a.firebaseapp.com",
    databaseURL: "https://mani-83d3a.firebaseio.com",
    projectId: "mani-83d3a",
    storageBucket: "mani-83d3a.appspot.com",
    messagingSenderId: "731257708840",
  };
  firebase.initializeApp(config);

  const messaging = firebase.messaging();

messaging.requestPermission().then(function() {
  console.log('Notification permission granted.');
  /*if(isTokenSentToServer()){
console.log('token alraedy saved');
  }else{
  	 getRegToken();
  }*/
  getRegToken();
  // TODO(developer): Retrieve an Instance ID token for use with FCM.
    
}).catch(function(err) {
  console.log('Unable to get permission to notify.', err);
});

function getRegToken(){

// Get Instance ID token. Initially this makes a network call, once retrieved
// subsequent calls to getToken will return from cache.
messaging.getToken().then(function(currentToken) {
  if (currentToken) {
    saveTokentoDatabase(currentToken);
    console.log(currentToken);
     setTokenSentToServer(true);
  } else {
    // Show permission request.
    console.log('No Instance ID token available. Request permission to generate one.');
    // Show permission UI.
    
    setTokenSentToServer(false);
  }
}).catch(function(err) {
  console.log('An error occurred while retrieving token. ', err);
  //showToken('Error retrieving Instance ID token. ', err);
  setTokenSentToServer(false);
});
}

function setTokenSentToServer(sent) {
    window.localStorage.setItem('sentToServer', sent ? 1 : 0);
  }
  

  /*function isTokenSentToServer() {
    return window.localStorage.getItem('sentToServer') === '1';
  }*/

   function saveTokentoDatabase(currentToken)
   {
$.ajax({

  url:'http://127.0.0.1:8000/insert',
  method:'POST',
  contentType: 'application/json',
  dataType: 'json',
  crossDomain: true,
  //data:currentToken
  data:'token='+JSON.stringify(currentToken)
}).done(function(result)
{
  console.log(result)
});

   }
</script>
</body>
</html>