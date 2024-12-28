import { getFileContent } from "../utils/tsutils";

type Bot = {
  pos: Coordinate;
  velocity: Coordinate;
}

type Coordinate = {
  x: number;
  y: number; 
}

const configuration = {
  seconds: 100,
  xLim: 101,
  yLim: 103
} as const;

// logic
const content = (await getFileContent('./input.txt')).split('\n')

const bots: Bot[] = content.map(line => {
  const coords = getBotFromLine(line)
  return {
    pos: { x: coords.x, y: coords.y },
    velocity: { x: coords.xVel, y: coords.yVel }
  }
});

for (let i = 0; i < configuration.seconds; i++){
  bots.forEach(bot => moveBot(bot, configuration.xLim, configuration.yLim));
}

const quadrants = countBotsInQuadrants(bots, configuration.xLim, configuration.yLim);

console.log(quadrants.q1 * quadrants.q2 * quadrants.q3 * quadrants.q4);
// functions

function moveBot(bot: Bot, xLim: number, yLim: number): void {
  let newX = bot.pos.x + bot.velocity.x;
  let newY = bot.pos.y + bot.velocity.y;

  if (newX < 0){
    newX = newX + xLim;
  } else if (newX >= xLim){
    newX = newX - xLim;
  }

  if (newY < 0){
    newY = newY + yLim;
  } else if (newY >= yLim){
    newY = newY - yLim;
  }

  bot.pos = { x: newX, y: newY };
}

function countBotsInQuadrants(bots: Bot[], xLim: number, yLim: number): { q1: number, q2: number, q3: number, q4: number } {

  const quadrants = {
    q1: 0,
    q2: 0,
    q3: 0,
    q4: 0
  }

  const mids = {
    x: Math.floor(xLim / 2),
    y: Math.floor(yLim / 2)
  }

  for (const bot of bots){
    const x = bot.pos.x;
    const y = bot.pos.y;

    // skip robots on the middle lines
    if (x == mids.x || y == mids.y){
      continue;
    }

    if (x > mids.x && y > mids.y){
      quadrants.q1++;
    } else if (x < mids.x && y > mids.y){
      quadrants.q2++;
    } else if (x < mids.x && y < mids.y){
      quadrants.q3++;
    } else if (x > mids.x && y < mids.y){
      quadrants.q4++;
    }
  }

  return quadrants;
}

function getBotFromLine(line: string) {
  
  // match all numbers
  const items = line.matchAll(/-?\d+/g);

  if (!items){
    throw new Error('Invalid input');
  }

  const [x, y, xVel, yVel] = Array.from(items).map(Number);
  return { x, y, xVel, yVel };
}

