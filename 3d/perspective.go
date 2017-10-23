package u3d


type Perspective struct {
	//	Whether this is a perspective-projection camera. Defaults to true.
	//	If false, no projection transformation is applied.
	Enabled bool

	//	Vertical field-of-view angle.
	FovY struct {
		//	In degrees. Defaults to 37.8493.
		Deg float64

		//	Deg-in-radians, times 0.5. This should always be kept in sync with Deg.
		RadHalf float64
	}

	//	Distance of the far-plane from the camera.
	ZFar float64

	//	Distance of the near-plane from the camera.
	ZNear float64
}
