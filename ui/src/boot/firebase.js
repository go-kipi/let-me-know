import { defineBoot } from 'quasar/wrappers'
import firebase from 'firebase/app'
import 'firebase/messaging' // For push notifications

export default defineBoot(({ app }) => {
  // Your Firebase project configuration
  const firebaseConfig = {
    apiKey: "YOUR_API_KEY",
    authDomain: "YOUR_PROJECT_ID.firebaseapp.com",
    databaseURL: "https://YOUR_PROJECT_ID.firebaseio.com",
    projectId: "YOUR_PROJECT_ID",
    storageBucket: "YOUR_PROJECT_ID.appspot.com",
    messagingSenderId: "YOUR_SENDER_ID",
    appId: "YOUR_APP_ID",
    measurementId: "YOUR_MEASUREMENT_ID"
  };

  // Initialize Firebase
  firebase.initializeApp(firebaseConfig);

  // Initialize Firebase Messaging
  const messaging = firebase.messaging();

  // Request permission for notifications
  messaging.requestPermission()
    .then(() => {
      console.log('Notification permission granted.');
      return messaging.getToken();
    })
    .then(token => {
      console.log('Firebase Token:', token);
      // Here, you can send the token to your backend to store it and send notifications later
    })
    .catch(err => {
      console.error('Error getting permission:', err);
    });

  // Handle foreground notifications
  messaging.onMessage(payload => {
    console.log('Message received. ', payload);
    // Handle the message and show a notification or update UI
  });

  app.config.globalProperties.$firebase = firebase;
});
