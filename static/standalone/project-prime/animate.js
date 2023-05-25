var counter = 0;

function startup() {
	setInterval('animate()',10);
}

function animate() {
	var is;
	var mult;
	var cells = document.getElementsByTagName('polygon');
	for (cell in cells) {
		if (cell >= 0) {
			is = cells[cell].getAttribute('rel');
			mult = Math.cos((1 - is) * counter * 0.05);
			cells[cell].setAttribute('fill-opacity',1- mult*mult*(1 - is));
		}
	}
	counter++;
}