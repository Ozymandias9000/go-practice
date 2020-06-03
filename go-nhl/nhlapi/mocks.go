package nhlApi

import (
	mock "github.com/jarcoal/httpmock"
)

var GetAllTeamsMock = mock.NewStringResponder(200, `{"teams" : [ {
    "id" : 1,
    "name" : "New Jersey Devils",
    "link" : "/api/v1/teams/1",
    "venue" : {
      "name" : "Prudential Center",
      "link" : "/api/v1/venues/null",
      "city" : "Newark",
      "timeZone" : {
        "id" : "America/New_York",
        "offset" : -4,
        "tz" : "EDT"
      }
    },
    "abbreviation" : "NJD",
    "teamName" : "Devils",
    "locationName" : "New Jersey",
    "firstYearOfPlay" : "1982",
    "division" : {
      "id" : 18,
      "name" : "Metropolitan",
      "nameShort" : "Metro",
      "link" : "/api/v1/divisions/18",
      "abbreviation" : "M"
    },
    "conference" : {
      "id" : 6,
      "name" : "Eastern",
      "link" : "/api/v1/conferences/6"
    },
    "franchise" : {
      "franchiseId" : 23,
      "teamName" : "Devils",
      "link" : "/api/v1/franchises/23"
    },
    "shortName" : "New Jersey",
    "officialSiteUrl" : "http://www.newjerseydevils.com/",
    "franchiseId" : 23,
    "active" : true
  }, {
    "id" : 2,
    "name" : "New York Islanders",
    "link" : "/api/v1/teams/2",
    "venue" : {
      "id" : 5026,
      "name" : "Barclays Center",
      "link" : "/api/v1/venues/5026",
      "city" : "Brooklyn",
      "timeZone" : {
        "id" : "America/New_York",
        "offset" : -4,
        "tz" : "EDT"
      }
    },
    "abbreviation" : "NYI",
    "teamName" : "Islanders",
    "locationName" : "New York",
    "firstYearOfPlay" : "1972",
    "division" : {
      "id" : 18,
      "name" : "Metropolitan",
      "nameShort" : "Metro",
      "link" : "/api/v1/divisions/18",
      "abbreviation" : "M"
    },
    "conference" : {
      "id" : 6,
      "name" : "Eastern",
      "link" : "/api/v1/conferences/6"
    },
    "franchise" : {
      "franchiseId" : 22,
      "teamName" : "Islanders",
      "link" : "/api/v1/franchises/22"
    },
    "shortName" : "NY Islanders",
    "officialSiteUrl" : "http://www.newyorkislanders.com/",
    "franchiseId" : 22,
    "active" : true
  }, {
    "id" : 3,
    "name" : "New York Rangers",
    "link" : "/api/v1/teams/3",
    "venue" : {
      "id" : 5054,
      "name" : "Madison Square Garden",
      "link" : "/api/v1/venues/5054",
      "city" : "New York",
      "timeZone" : {
        "id" : "America/New_York",
        "offset" : -4,
        "tz" : "EDT"
      }
    },
    "abbreviation" : "NYR",
    "teamName" : "Rangers",
    "locationName" : "New York",
    "firstYearOfPlay" : "1926",
    "division" : {
      "id" : 18,
      "name" : "Metropolitan",
      "nameShort" : "Metro",
      "link" : "/api/v1/divisions/18",
      "abbreviation" : "M"
    },
    "conference" : {
      "id" : 6,
      "name" : "Eastern",
      "link" : "/api/v1/conferences/6"
    },
    "franchise" : {
      "franchiseId" : 10,
      "teamName" : "Rangers",
      "link" : "/api/v1/franchises/10"
    },
    "shortName" : "NY Rangers",
    "officialSiteUrl" : "http://www.newyorkrangers.com/",
    "franchiseId" : 10,
    "active" : true
  }]}`)
