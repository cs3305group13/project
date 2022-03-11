
(() => {
    const main = document.querySelector('#main');
    if (!main.getContext) {
        return;
    }
    let ctx = main.getContext('2d');

    main.width = 1152;
    main.height = 648;

    const gravity = 0.5;

    class Player {
        constructor(standRight, standLeft, runRight, runLeft) {
            this.speed = 10;
            this.position = {
                x: 100,
                y: 100
            };
            this.width = 66;
            this.height = 150;
            this.velocity = {
                x: 0,
                y: 0
            };
            this.frames = 0;
            this.sprites = {
                stand: {
                    right: standRight,
                    cropWidth: 177,
                    width: 66
                },
                run: {
                    right: runRight,
                    cropWidth: 340,
                    width: 127.875
                }
            };
            this.currentSprite = this.sprites.stand.right;
            this.currentCropWidth = 177;
        }

        draw () {
            //ctx.fillStyle = 'green';
            //ctx.fillRect(this.position.x, this.position.y, this.width, this.height);
            ctx.drawImage(this.currentSprite, this.currentCropWidth * this.frames, 0, this.currentCropWidth, 400, this.position.x, this.position.y, this.width, this.height);
        }

        update () {
            this.frames++;
            if (this.frames > 59 && this.currentSprite === this.sprites.stand.right) {
                this.frames = 0
            } else if (this.frames > 29 && this.currentSprite === this.sprites.run.right) {
                this.frames = 0;
            }
            this.draw();
            this.position.x += this.velocity.x;
            this.position.y += this.velocity.y;
            if (this.position.y + this.height + this.velocity.y <= main.height) {
                this.velocity.y += gravity;
            }
        }
    }

    class Platform {
        constructor({x, y, image}) {
            this.position = {
                x,
                y
            }
            this.image = image;
            this.width = image.width;
            this.height = image.height;
        }

        draw() {
            ctx.drawImage(this.image, this.position.x, this.position.y);
        }
    }

    class GenericObject {
        constructor({x, y, image}) {
            this.position = {
                x,
                y
            }
            this.image = image;
            this.width = image.width;
            this.height = image.height;
        }

        draw() {
            ctx.drawImage(this.image, this.position.x, this.position.y);
        }
    }

    function createImage(imageSrc) {
        const image = new Image();
        image.src = imageSrc;
        return image;
    }

    let standRightImage = createImage("../media/sideScrollerAssets/spriteStandRight.png");
    let standLeftImage = createImage("../media/sideScrollerAssets/spriteStandLeft.png");
    let runRightImage = createImage("../media/sideScrollerAssets/spriteRunRight.png");
    let runLeftImage = createImage("../media/sideScrollerAssets/spriteRunLeft.png");

    let platformImage = createImage("../media/sideScrollerAssets/platform.png");
    let backgroundImage = createImage("../media/sideScrollerAssets/background.png");
    let hillImage = createImage("../media/sideScrollerAssets/hills.png");
    //let diamondImage = createImage("../media/sideScrollerAssets/diamond.png");


    let player = new Player(standRightImage, standLeftImage, runRightImage, runLeftImage);

    let platforms = [new Platform({
                           x: 0, y: 523, image: platformImage
                       }),
                       new Platform({
                           x: platformImage.width + 200, y: 523, image: platformImage
                       }),
        new Platform({
            x: (platformImage.width * 2) + 500, y: 523, image: platformImage}),
        new Platform({
            x: (platformImage.width * 3) + 700, y: 300, image: platformImage}),
        new Platform({
            x: (platformImage.width * 4) + 1000, y: 400, image: platformImage}),
        new Platform({
            x: (platformImage.width * 5) + 1300, y: 200, image: platformImage})
    ];

    let genericObjects = [new GenericObject({x: 0, y: 0, image: backgroundImage}),
                            new GenericObject({x: 0, y: 0, image: hillImage})
        /*new GenericObject({x: (580 * 5) + 1400, y: 150, image: diamondImage})*/];

    let keys = {
        right: {
            pressed: false
        },
        left: {
            pressed: false
        }
    }

    let scrollOffset = 0;
    
    function init() {
        standRightImage = createImage("../media/sideScrollerAssets/spriteStandRight.png");
        standLeftImage = createImage("../media/sideScrollerAssets/spriteStandLeft.png");
        runRightImage = createImage("../media/sideScrollerAssets/spriteRunRight.png");
        runLeftImage = createImage("../media/sideScrollerAssets/spriteRunLeft.png");

        standRightImage = createImage("../media/sideScrollerAssets/spriteStandRight.png");
        platformImage = createImage("../media/sideScrollerAssets/platform.png");
        backgroundImage = createImage("../media/sideScrollerAssets/background.png");
        hillImage = createImage("../media/sideScrollerAssets/hills.png");
        //diamondImage = createImage("../media/sideScrollerAssets/diamond.png");

        player = new Player(standRightImage, standLeftImage, runRightImage, runLeftImage);

        platforms = [new Platform({
            x: 0, y: 523, image: platformImage
        }),
            new Platform({
                x: platformImage.width + 200, y: 523, image: platformImage
            }),
            new Platform({
                x: (platformImage.width * 2) + 500, y: 523, image: platformImage}),
            new Platform({
                x: (platformImage.width * 3) + 700, y: 300, image: platformImage}),
            new Platform({
                x: (platformImage.width * 4) + 1000, y: 400, image: platformImage}),
            new Platform({
                x: (platformImage.width * 5) + 1300, y: 200, image: platformImage})
        ];

        genericObjects = [new GenericObject({x: 0, y: 0, image: backgroundImage}),
            new GenericObject({x: 0, y: 0, image: hillImage})];

        keys = {
            right: {
                pressed: false
            },
            left: {
                pressed: false
            }
        }

        scrollOffset = 0;
    }

    function animate() {
        requestAnimationFrame(animate);
        ctx.clearRect(0, 0, main.width, main.height);

        genericObjects.forEach(genericObject => {
            genericObject.draw();
        });

        platforms.forEach(platform => {
            platform.draw();
        });

        player.update();

        if (keys.right.pressed && player.position.x < 400) {
            player.velocity.x = player.speed;
        } else if ((keys.left.pressed && player.position.x > 20) ||
                   (keys.left.pressed && scrollOffset === 0 && player.position.x > 0)) {
            player.velocity.x = -player.speed;
        } else {
            player.velocity.x = 0;
            if (keys.right.pressed) {
                scrollOffset += 5;
                platforms.forEach(platform => {
                    platform.position.x -= player.speed;
                });
                genericObjects.forEach(genericObject => {
                    genericObject.position.x -= player.speed * 0.66;
                })
            } else if (keys.left.pressed && scrollOffset > 0) {
                scrollOffset -= 5;
                platforms.forEach(platform => {
                    platform.position.x += player.speed;
                });
                genericObjects.forEach(genericObject => {
                    genericObject.position.x += player.speed * 0.66;
                })
            }
        }

        platforms.forEach(platform => {
            if (player.position.y + player.height <= platform.position.y &&
                player.position.y + player.height +player.velocity.y >= platform.position.y &&
                player.position.x + player.width >= platform.position.x &&
                player.position.x <= platform.position.x + platform.width) {
                player.velocity.y = 0;
            }
        });
        if (scrollOffset > 2000) {
            console.log('you win!');
            alert('you win! Score 1274 Level 1 Cleared')
        }

        if (player.position.y > main.height) {
            init();
        }
    }

    animate();

    window.addEventListener('keydown', ({ keyCode }) => {
        switch (keyCode) {
            case 65:
                console.log('left');
                keys.left.pressed = true;
                break;

            case 83:
                console.log('down');
                break;

            case 68:
                console.log('right');
                keys.right.pressed = true;
                player.currentSprite = player.sprites.run.right;
                player.currentCropWidth = player.sprites.run.cropWidth;
                player.width = player.sprites.run.width;
                break;

            case 87:
                console.log('up');
                player.velocity.y -= 15;
                break;
        }
    });

    window.addEventListener('keyup', ({ keyCode }) => {
        switch (keyCode) {
            case 65:
                console.log('left');
                keys.left.pressed = false;
                break;

            case 83:
                console.log('down');
                break;

            case 68:
                console.log('right');
                keys.right.pressed = false;
                player.currentSprite = player.sprites.stand.right;
                player.currentCropWidth = player.sprites.stand.cropWidth;
                player.width = player.sprites.stand.width;
                break;

            case 87:
                console.log('up');
                break;
        }
    });

})();