(function(){

  var keyConfig;

  //private

  function send(body){
    var URL = 'http://localhost:6969/v1/notification';
    var METHOD = 'post';
    var request = new XMLHttpRequest();

    request.open(METHOD, URL);
    request.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    request.send(body);
  }

  function buildNotification(title, tag1, tag2, tag3, info, relevance, type, isTransaction, transaction, transactionResult, company){
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
    if(tags.constructor === Array){
      return buildNotification(title, tags[0], tags[1], tags[2], info, relevance, 1, false, null, null, keyConfig);
    }
    return buildNotification(title, null, null, null, info, relevance, 1, false, null, null, keyConfig);
  }

  function fillTransaction(name, info, tags, scenario){
    if(tags.constructor === Array){
      return buildNotification(null, tags[0], tags[1], tags[2], info, 2, 2, true, name, scenario, keyConfig);
    }
    return buildNotification(null, null, null, null, info, 2, 2, true, name, scenario, keyConfig);
  }

  //exports

  var factory = {};

  factory.config = function(key){
    keyConfig = key;
  };

  //notify

  factory.notify = {};

  factory.notify.log = function(title, info, relevance, tags){
    send(
      fillNotify(title, info, relevance, tags, 1)
    );
  };

  factory.notify.info = function(title, info, relevance, tags){
    send(
      fillNotify(title, info, relevance, tags, 2)
    );
  };

  factory.notify.warning = function(title, info, relevance, tags){
    send(
      fillNotify(title, info, relevance, tags, 3)
    );
  };

  factory.notify.error = function(title, info, relevance, tags){
    send(
      fillNotify(title, info, relevance, tags, 4)
    );
  };

  //transaction

  factory.transaction = {};

  factory.transaction.begin = function(name, info, tags){
    send(
      fillTransaction(name, info, tags, false)
    );
  };

  factory.transaction.end = function(name, info, tags){
    send(
      fillTransaction(name, info, tags, true)
    );
  };

  factory.transaction.fail = function(name, info, tags){
    send(
      fillTransaction(name, info, tags, false)
    );
  };

  window.go = factory;
})();
