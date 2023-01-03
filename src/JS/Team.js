export default class Team {

    name;
    gamesPlayed;
    draws;
    wins;
    losses;
    points;

    constructor(name, gamesPlayed, draws, wins, losses, points) {
        this.name = name;
        this.gamesPlayed = gamesPlayed;
        this.draws = draws;
        this.wins = wins;
        this.losses = losses;
        this.points = points;
    }      

    
    static generateTeams() {
        
        const teams = ["Liverpool", "Manchester United", "Manchester City", "Everton", "Newcastle", "Aston Villa", "Southhampton", "Leicster", "Arsenal", "Tottenham Hotspurs", "Brighton", "Wolves", "Chelsea", "Brentford", "Fulham", "Leeds", "Crystal Palace", "Bournemouth", "West Ham", "Wolves", "Nottm Forrest"];
        
        const teamList = [];

        teams.forEach(team => {
            const numOfDraws = Math.random(5);
            const numOfWins = Math.random(11);
            const numOfLosses = 15 - numOfDraws - numOfLosses;
            const points = numOfDraws + (numOfWins*3);
            const newTeam = new Team(team, 15, numOfDraws, numOfWins, numOfLosses, points)
            teamList.push(newTeam)
        });

        for (let index = 0; index < teamList.length; index++) {
            if(index < teamList.length-1 && teamList[index].points > teamList[index+1].points) {
                const temp = teamList[index];
                teamList[index] = teamList[index+1];
                teamList[index+1] = temp;
                index = 0;
            }
        }

        return teamList;
    }

    
}
