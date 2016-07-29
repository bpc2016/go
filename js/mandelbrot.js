$(document).ready(function(){
	console.log('ok ...')
	var ind = 1;
	// fetch the image with an ajax call ..
	function fetchPiece() {
   	  $.get('http://localhost:8000/image')
   	    .done(function(result){
   		console.log('got: '+ result ? 'piece z=' : 0, ind, ' starts: ', "'"+result.substr(0,30)+"'")
		if (result==='0'){
			console.log('terminate!');
			return;
		} 
   		if (result) {
			var h = ['<img ',
			'style="position:absolute; top:0; left:0; z-index:',
			 ind,
			';" src ="data:image/png;base64,',
			result,
   			'"></img>'].join('');
   			$(h).appendTo("#imgs");
   			ind+=2;
			fetchPiece()
   		} else {
   			console.log('no more pieces')
   		}
   	   })
   	   .fail(function(){
   		console.log('oops!')
   	   });
	}
	fetchPiece()
})
