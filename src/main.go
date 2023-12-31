package main

import rl "github.com/gen2brain/raylib-go/raylib"

//-----------------------------------------------------
var running bool=true;
const w,h=640,480;
var bgcol = rl.NewColor(111,02,244,255);
var plSprite rl.Texture2D;
var plRect=rl.NewRectangle(50,50,32,32);
var pspeed=5.5;
var music rl.Music=rl.LoadMusicStream("res/mus.mp3");
var in = make(map[string]bool);
var cam rl.Camera2D;
var t int64;
func recCenter(r rl.Rectangle) rl.Vector2{
   return rl.NewVector2(r.X+r.Width, r.Y+r.Height);
}
//-----------------------------------------------------
type ball struct{
   pos rl.Vector2;
   vel rl.Vector2;
   rec rl.Rectangle;
   col rl.Color;
}
func newBall(p rl.Vector2) ball{
   var b = ball{pos: p,vel: rl.NewVector2(1,1),rec: rl.NewRectangle(16,16,32,32), col: rl.Red };
   return b;

}
func moveBall(b ball,dt float32){
   b.pos.X+=b.vel.X*dt;
   b.pos.Y+=b.vel.Y*dt; 
}
func drawball(b ball){
   rl.DrawTextureRec(plSprite, rl.NewRectangle(64*float32(t/3),0,64,64),b.pos, b.col);
}
//-Stuff-----------------------------------------------
func event(){
   in["up"]=rl.IsKeyDown(rl.KeyW);
   in["down"]=rl.IsKeyDown(rl.KeyS);
   in["left"]=rl.IsKeyDown(rl.KeyA);
   in["right"]=rl.IsKeyDown(rl.KeyD);
}
func tick(dt float64){
   dt*=60;
   if(in["up"]){plRect.Y-=float32(pspeed*dt);}
   if(in["down"]){plRect.Y+=float32(pspeed*dt);}
   if(in["left"]){plRect.X-=float32(pspeed*dt);}
   if(in["right"]){plRect.X+=float32(pspeed*dt);}
}
func render(){
   cam.Target=recCenter(plRect);
   rl.BeginDrawing();
   rl.BeginMode2D(cam)
   rl.ClearBackground(bgcol);
   
   rl.DrawTexture(plSprite, int32(0), int32(0), rl.White);
   //rl.DrawTexture(plSprite, int32(plRect.X), int32(plRect.Y), rl.White);
   rl.DrawTextureRec(plSprite, rl.NewRectangle(64*float32(t/3),0,64,64),rl.NewVector2((plRect.X), (plRect.Y)), rl.White);
   
   rl.EndMode2D();
   rl.EndDrawing();
}
//-Oneers----------------------------------------------
func gameLoop(){
   event();
   tick(1.0/60);
   render();
   t++;
   if(rl.WindowShouldClose()){running=false;}
   if(running){gameLoop();}
}
func initalize(){
	rl.InitWindow(w, h, "African Daisy");
  // rl.SetExitKey(0); make escape not close 
   rl.SetTargetFPS(60);
   plSprite=rl.LoadTexture("res/cube.png")
   rl.InitAudioDevice()
   cam = rl.NewCamera2D(rl.NewVector2(w/2,h/2),rl.NewVector2(0,0),0,3);
   rl.PlayMusicStream(music);
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
