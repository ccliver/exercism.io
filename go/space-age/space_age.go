package space

// Planet defines a planet in the solar system.
type Planet string

const earthOrbitalPeriod float64 = 31557600

// Age calculates your age in seconds on another planet in the solar system.
func Age(seconds float64, planet Planet) float64 {
  orbitalPeriodSeconds := map[Planet]float64 {
    "Earth": earthOrbitalPeriod,
    "Mercury": earthOrbitalPeriod * 0.2408467,
    "Venus": earthOrbitalPeriod * 0.61519726,
    "Mars": earthOrbitalPeriod * 1.8808158,
    "Jupiter": earthOrbitalPeriod * 11.862615,
    "Saturn": earthOrbitalPeriod * 29.447498,
    "Uranus": earthOrbitalPeriod * 84.016846,
    "Neptune": earthOrbitalPeriod * 164.79132,
  }

  return seconds / orbitalPeriodSeconds[planet]
}
