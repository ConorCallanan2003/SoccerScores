import Team from './Team.js';
export default class League {

    name;
    teams;

    constructor(name) {
        this.name = name;
        this.teams = [];
    }

    static retrieveLeagues() {

        const leagues = [];

        var xmlHttp = new XMLHttpRequest();
        xmlHttp.open( "GET", "http://52.214.158.151:3333/leagues", false ); // false for synchronous request
        xmlHttp.send( null );

        const obj = JSON.parse(xmlHttp.responseText);
        
        obj.forEach(element => {
            const newLeague = new League(element.Name);
            
            element.Teams.forEach(team => {
                const newTeam = new Team(team.Team_Name, team.Overall.Games_Played, team.Overall.Draw, team.Overall.Won, team.Overall.Lost, team.Overall.Points);
                newLeague.teams.push(newTeam);
            });
    
            leagues.push(newLeague)    
        });

        console.log(leagues);
        return leagues;

    }

}
