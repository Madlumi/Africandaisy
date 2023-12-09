package main
import rl "github.com/gen2brain/raylib-go/raylib"

var running bool=true;
const w,h=640,480;
var bgcol = rl.NewColor(111,02,244,255);
var plSprite rl.Texture2D;


func main() {
   initalize();
   gameLoop();
   done();
}
func initalize(){
	rl.InitWindow(w, h, "African Daisy");
   rl.SetExitKey(0);
   rl.SetTargetFPS(60);
   plSprite=rl.LoadTexture("res/cube.png")
}
func gameLoop(){
   for (running){
      event();
      tick();
      render();
      if (rl.WindowShouldClose()){running=false;}
   }
}
func event(){}
func tick(){}
func render(){
		rl.BeginDrawing();
		rl.ClearBackground(bgcol);
      rl.DrawTexture(plSprite, 50, 50, rl.White);
		rl.EndDrawing();
}
func done(){
	rl.CloseWindow();
}
