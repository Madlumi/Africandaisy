package main

import rl "github.com/gen2brain/raylib-go/raylib"

//-----------------------------------------------------
var running bool=true;
const w,h=640,480;
var bgcol = rl.NewColor(111,02,244,255);
var plSprite rl.Texture2D;
var px,py int32=50,50;
var in = make(map[string]bool);
//-Stuff-----------------------------------------------
func event(){
   in["up"]=rl.IsKeyDown(rl.KeyW);
   in["down"]=rl.IsKeyDown(rl.KeyS);
   in["left"]=rl.IsKeyDown(rl.KeyA);
   in["right"]=rl.IsKeyDown(rl.KeyD);
}
func tick(){
   if(in["up"]){py--;}
   if(in["down"]){py++;}
   if(in["left"]){px--;}
   if(in["right"]){px++;}
}
func render(){
   rl.BeginDrawing();
   rl.ClearBackground(bgcol);
   rl.DrawTexture(plSprite, px, py, rl.White);
   rl.EndDrawing();
}
//-Oneers----------------------------------------------
func gameLoop(){
   event();
   tick();
   render();
   if(rl.WindowShouldClose()){running=false;}
   if(running){gameLoop();}
}
func initalize(){
	rl.InitWindow(w, h, "African Daisy");
  // rl.SetExitKey(0); make escape not close 
   rl.SetTargetFPS(60);
   plSprite=rl.LoadTexture("res/cube.png")
}
func done(){
   rl.UnloadTexture(plSprite);
	rl.CloseWindow();
}
func main() {
   initalize();
   gameLoop();
   done();
}
