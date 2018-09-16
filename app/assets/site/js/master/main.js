var $JSModel

//initTxJsonObj initialize json string from golang to json object
//arg
//  jsonString string
function initTxJsonObj(jsonString){
    $JSModel = JSON.parse(jsonString)
}