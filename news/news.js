var page = require('webpage').create();
page.open('url_of_news_site', function(status) {
  if (status !== 'success') {
    console.log('Unable to access network');
  } else {
    var body = page.evaluate(function() {
      return document.getElementsByTagName('body')[0].innerHTML
    });
    console.log(body);
  }
  phantom.exit();
});
