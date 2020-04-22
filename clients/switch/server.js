const path = require('path');


const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const PROTO_PATH = path.join(__dirname, '..', '..', 'framebuffer', 'framebuffer.proto');
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const framebuffer = grpc.loadPackageDefinition(packageDefinition);

const client = new framebuffer.Drawer('raspi.cacti.makerforce.io:5000', grpc.credentials.createInsecure());


const express = require('express');

const app = express();
const PORT = 5001;

const PIXEL_COUNT = 275;

// TODO: async
app.post('/on', (req, res, next) => {
	let fill = 0x00000000;
	switch (req.query.tempreature) {
	case 'cool':
		fill = 0xFFFFFFFF;
		break;
	case 'warm':
	default:
		fill = 0xFFFFAA77;
		break;
	}
	client.drawFrame({
		frame: {
			dots: Array(PIXEL_COUNT-1).fill(fill).concat([0xBBFF0000]),
		},
		layer: 'LIGHT',
	}, (err, resp) => {
		if (err) {
			return next(err);
		}
		res.writeHead(200);
		res.end();
	});
});
app.post('/off', (req, res, next) => {
	client.drawFrame({
		frame: {
			dots: Array(PIXEL_COUNT).fill(0x00000000),
		},
		layer: 'LIGHT',
	}, (err, resp) => {
		if (err) {
			return next(err);
		}
		res.writeHead(200);
		res.end();
	});
});

app.use(express.static('public'));

app.listen(PORT, () => console.log(`App listening on ${PORT}`));
