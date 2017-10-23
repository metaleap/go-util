# u3d

Spatial data types & helpers for use in 3D apps (AABB, Frustum, bounding volumes etc)

## Usage

#### type AaBb

```go
type AaBb struct {
	Min, Max, Center, Extent unum.Vec3
}
```


#### func (*AaBb) BoundingSphere

```go
func (me *AaBb) BoundingSphere(center *unum.Vec3) (radius float64)
```

#### func (*AaBb) Clear

```go
func (me *AaBb) Clear()
```

#### func (*AaBb) ResetMinMax

```go
func (me *AaBb) ResetMinMax()
```

#### func (*AaBb) SetCenterExtent

```go
func (me *AaBb) SetCenterExtent()
```

#### func (*AaBb) SetMinMax

```go
func (me *AaBb) SetMinMax()
```

#### func (*AaBb) Transform

```go
func (me *AaBb) Transform(mat *unum.Mat4)
```

#### func (*AaBb) UpdateMinMax

```go
func (me *AaBb) UpdateMinMax(vec *unum.Vec3)
```

#### func (*AaBb) UpdateMinMaxFrom

```go
func (me *AaBb) UpdateMinMaxFrom(aabb *AaBb)
```

#### type Bounds

```go
type Bounds struct {
	Sphere float64
	AaBox  AaBb
}
```


#### func (*Bounds) Clear

```go
func (me *Bounds) Clear()
```

#### func (*Bounds) Reset

```go
func (me *Bounds) Reset()
```

#### type Frustum

```go
type Frustum struct {
	Bounding Bounds
	Planes   [6]FrustumPlane
	Axes     struct {
		X, Y, Z unum.Vec3
	}
	Near, Far FrustumCoords
}
```


#### func (*Frustum) HasPoint

```go
func (me *Frustum) HasPoint(pos, point *unum.Vec3, zNear, zFar float64) bool
```

#### func (*Frustum) HasSphere

```go
func (me *Frustum) HasSphere(pos, center *unum.Vec3, radius, zNear, zFar float64) (fullyInside, intersect bool)
```

#### func (*Frustum) UpdateAxes

```go
func (me *Frustum) UpdateAxes(dir, upVector, upAxis *unum.Vec3)
```

#### func (*Frustum) UpdateAxesCoordsPlanes

```go
func (me *Frustum) UpdateAxesCoordsPlanes(persp *Perspective, pos, dir, upVector, upAxis *unum.Vec3)
```

#### func (*Frustum) UpdateCoords

```go
func (me *Frustum) UpdateCoords(persp *Perspective, pos *unum.Vec3)
```

#### func (*Frustum) UpdatePlanes

```go
func (me *Frustum) UpdatePlanes()
```

#### func (*Frustum) UpdatePlanesGH

```go
func (me *Frustum) UpdatePlanesGH(mat *unum.Mat4, normalize bool)
```
Gribb/Hartmann: "Fast Extraction of Viewing Frustum Planes from the
WorldView-Projection Matrix"

#### func (*Frustum) UpdateRatio

```go
func (me *Frustum) UpdateRatio(persp *Perspective, aspectRatio float64)
```

#### type FrustumCoords

```go
type FrustumCoords struct {
	unum.Vec2
	C, TL, TR, BL, BR unum.Vec3
}
```


#### type FrustumPlane

```go
type FrustumPlane struct {
	unum.Vec4
}
```


#### func (*FrustumPlane) Normalize

```go
func (me *FrustumPlane) Normalize()
```

#### type MeshDescF3

```go
type MeshDescF3 struct {
	//	The indexed vertices making up this triangle face.
	V [3]MeshDescF3V

	//	ID, Tags
	MeshFaceBase
}
```

Represents an indexed triangle face.

#### func  NewMeshDescF3

```go
func NewMeshDescF3(tags, id string, verts ...MeshDescF3V) (me *MeshDescF3)
```
Creates and initializes a new MeshDescF3V with the specified tags, ID and verts,
and returns it. tags may be empty or contain multiple classification tags
separated by spaces, which will be split into Tags.

#### type MeshDescF3V

```go
type MeshDescF3V struct {
	//	Index of the vertex position
	PosIndex uint32

	//	Index of the texture-coordinate.
	TexCoordIndex uint32

	//	Index of the vertex normal.
	NormalIndex uint32
}
```

Represents an indexed vertex in a MeshDescF3.

#### type MeshDescVA2

```go
type MeshDescVA2 [2]float32
```

Represents a 2-component vertex attribute in a MeshDescriptor. (such as for
example texture-coordinates)

#### type MeshDescVA3

```go
type MeshDescVA3 [3]float32
```

Represents a 3-component vertex attribute in a MeshDescriptor (such as for
example vertex-normals)

#### func (*MeshDescVA3) ToVec3

```go
func (me *MeshDescVA3) ToVec3(vec *unum.Vec3)
```

#### type MeshDescriptor

```go
type MeshDescriptor struct {
	//	Vertex positions
	Positions []MeshDescVA3

	//	Vertex texture coordinates
	TexCoords []MeshDescVA2

	//	Vertex normals
	Normals []MeshDescVA3

	//	Indexed triangle definitions
	Faces []MeshDescF3
}
```

Represents yet-unprocessed, descriptive mesh source data.

#### func  MeshDescriptorCube

```go
func MeshDescriptorCube() (meshDescriptor *MeshDescriptor, err error)
```
A MeshProvider that creates MeshDescriptor for a cube with extents -1 .. 1. args
is ignored and err is always nil. The returned MeshDescriptor contains 12
triangle faces with IDs "t0" through "t11". These faces are classified in 6
distinct tags: "front","back","top","bottom","right","left".

#### func  MeshDescriptorPlane

```go
func MeshDescriptorPlane() (meshDescriptor *MeshDescriptor, err error)
```
A MeshProvider that creates MeshDescriptor for a flat ground plane with extents
-1 .. 1. args is ignored and err is always nil. The returned MeshDescriptor
contains 2 triangle faces with IDs "t0" through "t1". These faces are all
classified with tag: "plane".

#### func  MeshDescriptorPyramid

```go
func MeshDescriptorPyramid() (meshDescriptor *MeshDescriptor, err error)
```
A MeshProvider that creates MeshDescriptor for a pyramid with extents -1 .. 1.
args is ignored and err is always nil. The returned MeshDescriptor contains 4
triangle faces with IDs "t0" through "t3". These faces are all classified with
tag: "pyr".

#### func  MeshDescriptorQuad

```go
func MeshDescriptorQuad() (meshDescriptor *MeshDescriptor, err error)
```
A MeshProvider that creates MeshDescriptor for a quad with extents -1 .. 1. args
is ignored and err is always nil. The returned MeshDescriptor contains 2
triangle faces with IDs "t0" through "t1". These faces are all classified with
tag: "quad".

#### func  MeshDescriptorTri

```go
func MeshDescriptorTri() (meshDescriptor *MeshDescriptor, err error)
```
A MeshProvider that creates MeshDescriptor for a triangle with extents -1 .. 1.
args is ignored and err is always nil. The returned MeshDescriptor contains 1
triangle face with ID "t0" and tag "tri".

#### func (*MeshDescriptor) AddFaces

```go
func (me *MeshDescriptor) AddFaces(faces ...*MeshDescF3)
```
Adds all specified Faces to this MeshDescriptor.

#### func (*MeshDescriptor) AddNormals

```go
func (me *MeshDescriptor) AddNormals(normals ...MeshDescVA3)
```
Adds all the specified Normals to this MeshDescriptor.

#### func (*MeshDescriptor) AddPositions

```go
func (me *MeshDescriptor) AddPositions(positions ...MeshDescVA3)
```
Adds all specified Positions to this MeshDescriptor.

#### func (*MeshDescriptor) AddTexCoords

```go
func (me *MeshDescriptor) AddTexCoords(texCoords ...MeshDescVA2)
```
Adds all the specified TexCoords to this MeshDescriptor.

#### type MeshFaceBase

```go
type MeshFaceBase struct {
	//	Mesh-unique identifier for this face.
	ID string

	//	Arbitrary classification tags for this face.
	Tags []string
}
```


#### type MeshProvider

```go
type MeshProvider func() (*MeshDescriptor, error)
```


#### type Perspective

```go
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
```
