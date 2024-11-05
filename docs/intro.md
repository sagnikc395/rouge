## Rougelike Intro 

- Usign Ebiten engine to make my first game.
- A simple ECS to make our application data driven and allow for much greater flexibility later on.
  
1. Creating a Game Struct
   1. Hold all the data we need globally for the game and also will be the structure which will meet the interface to bootstrap Ebiten.
2. Define the constructor for the Game struct :- NewGame 
3. Add the callbacks for Update, Draw and Layout:
   1. func (g *Game) Update() error {}
   2. func (g *Game) Draw(screen *ebiten.Image) {}
   3. func (g *Game) Layout(w,h int)(int,int) {}
   4. here we are just telling Ebiten to create a window of 1280 * 800 .
4. define the main method and set the things.
    