(function(){

  var keyConfig;
  var options = {};
  var URL = "http://localhost:6969/v1/notification";

  //private

  function send(body){
    var METHOD = 'post';
    var request = new XMLHttpRequest();

    request.open(METHOD, URL);
    request.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    request.send(body);
  }

  function buildModel(title, tag1, tag2, tag3, info, relevance, type, isTransaction, transaction, transactionResult, company){
    var model = {
      title: title,
      tag1: tag1,
      tag2: tag2,
      tag3: tag3,
      info: info,
      relevance: relevance,
      type: type,
      isTransaction: isTransaction,
      transaction: transaction,
      transactionResult: transactionResult,
      company: company
    };

    return JSON.stringify(model);
  }

  function fillNotify(title, info, relevance, tags, type){
    if(!!tags && tags.constructor === Array){
      return buildModel(title, tags[0], tags[1], tags[2], info, relevance, 1, false, null, null, keyConfig);
    }
    return buildModel(title, null, null, null, info, relevance, 1, false, null, null, keyConfig);
  }

  function fillTransaction(name, info, tags, scenario){
    if(!!tags && tags.constructor === Array){
      return buildModel(null, tags[0], tags[1], tags[2], info, 2, 2, true, name, scenario, keyConfig);
    }
    return buildModel(null, null, null, null, info, 2, 2, true, name, scenario, keyConfig);
  }

  function doOutputNotify(title, info, relevance, tags, type){
    switch(options.output){
      case 'console':
        console[type](fillNotify(title, info, relevance, tags, 1));
        break;
      case 'api':
        send(fillNotify(title, info, relevance, tags, 1));
        break;
      default:
        send(fillNotify(title, info, relevance, tags, 1));
    }
  }

  function doOutputTransaction(name, info, tags, type){
    switch(options.output){
      case 'console':
        console.info(fillTransaction(name, info, tags, type));
        break;
      case 'api':
        send(fillTransaction(name, info, tags, type));
        break;
      default:
        send(fillTransaction(name, info, tags, type));
    }
  }

  //exports

  var factory = {};

  factory.config = function(key, opts){
    keyConfig = key;
    options.output = opts.output || 'api';

    //TODO if success
    kidnapXHR(XMLHttpRequest);
  };

  //notify

  factory.notify = {};

  factory.notify.log = function(title, info, relevance, tags){
    doOutputNotify(title, info, relevance, tags, 'log');
  };

  factory.notify.info = function(title, info, relevance, tags){
    doOutputNotify(title, info, relevance, tags, 'info');
  };

  factory.notify.warning = function(title, info, relevance, tags){
    doOutputNotify(title, info, relevance, tags, 'warn');
  };

  factory.notify.error = function(title, info, relevance, tags){
    doOutputNotify(title, info, relevance, tags, 'error');
  };

  //transaction

  factory.transaction = {};

  factory.transaction.begin = function(name, info, tags){
    doOutputTransaction(name, info, tags, false);
  };

  factory.transaction.end = function(name, info, tags){
    doOutputTransaction(name, info, tags, true);
  };

  factory.transaction.fail = function(name, info, tags){
    doOutputTransaction(name, info, tags, false);
  };

  function kidnapXHR(XHR){
    var open = XHR.prototype.open;
    var send = XHR.prototype.send;
    var timeBegin;

    XHR.prototype.open = function(method, url, async, user, pass) {
        this._url = url;
        open.call(this, method, url, async, user, pass);
    };

    XHR.prototype.send = function(data) {
        var self = this;
        var oldOnReadyStateChange;
        var url = this._url;

        function onReadyStateChange() {
            if(self.readyState == 4) {
              if(this.responseURL !== URL){

                var info = {
                  length: this.response.length * 16 + ' bytes',
                  status: this.status,
                  url: this.responseURL,
                  delay: (Date.now() - timeBegin) + ' duration'
                };

                switch(this.status){
                  case 200:
                    factory.notify.log('request http', JSON.stringify(info), 1, null);
                    break;
                  case 304:
                    factory.notify.log('request http', JSON.stringify(info), 1, null);
                    break;
                  case 404:
                    factory.notify.error('request http', JSON.stringify(info), 2, null);
                    break;
                  case 500:
                    factory.notify.error('request http', JSON.stringify(info), 3, null);
                    break;
                }
              }
            }

            if(oldOnReadyStateChange) {
                oldOnReadyStateChange();
            }
        }

        if(!this.noIntercept) {
            if(this.addEventListener) {
                this.addEventListener("readystatechange", onReadyStateChange, false);
            } else {
                oldOnReadyStateChange = this.onreadystatechange;
                this.onreadystatechange = onReadyStateChange;
            }
        }

    timeBegin = Date.now();
        send.call(this, data);
    };
  }

  window.go = factory;
})();
