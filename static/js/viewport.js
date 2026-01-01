/* jshint esversion: 6 */

import * as THREE from 'three';

const container = document.getElementById('canvas-container');
const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(75, container.clientWidth / 500, 0.1, 1000);
const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });

renderer.setSize(container.clientWidth, 500);
container.appendChild(renderer.domElement);

// Crear Pirámide Dodecagonal
const geometry = new THREE.ConeGeometry(10, 30, 12);
const material = new THREE.MeshNormalMaterial({ wireframe: true });
const pyramid = new THREE.Mesh(geometry, material);
scene.add(pyramid);

camera.position.z = 40;
camera.position.y = 30;

let falling = false;
let speed = 0;

document.getElementById('start-fall').onclick = () => {
    falling = true;
    camera.position.y = 30; // Reset posición
    camera.fov = 75;        // Reset FOV
};

function animate() {
    requestAnimationFrame(animate);
    
    if (falling && camera.position.y > 0) {
        speed += 0.01;
        camera.position.y -= speed;
        // Efecto Vértigo (Aumento de FOV)
        camera.fov += 0.5;
        camera.updateProjectionMatrix();
    } else if (camera.position.y <= 0) {
        falling = false;
        speed = 0;
    }

    pyramid.rotation.y += 0.005;
    renderer.render(scene, camera);
}

animate();