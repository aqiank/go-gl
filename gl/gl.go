package gl

//#cgo LDFLAGS: -lGL
//#include <GL/gl.h>
//#include <string.h>
import "C"
import "unsafe"
import "reflect"

const (
	VERSION_1_1 = 1
	VERSION_1_2 = 1
	VERSION_1_3 = 1
	ARB_imaging = 1

	/* Boolean values */
	FALSE =					0
	TRUE =					1

	/* Data types */
	BYTE =					0x1400
	UNSIGNED_BYTE =				0x1401
	SHORT =					0x1402
	UNSIGNED_SHORT =			0x1403
	INT =					0x1404
	UNSIGNED_INT =				0x1405
	FLOAT =					0x1406
	_2_BYTES =				0x1407
	_3_BYTES =				0x1408
	_4_BYTES =				0x1409
	DOUBLE =				0x140A

	/* Primitives */
	POINTS =				0x0000
	LINES =					0x0001
	LINE_LOOP =				0x0002
	LINE_STRIP =				0x0003
	TRIANGLES =				0x0004
	TRIANGLE_STRIP =			0x0005
	TRIANGLE_FAN =				0x0006
	QUADS =					0x0007
	QUAD_STRIP =				0x0008
	POLYGON	=				0x0009

	/* Vertex Arrays */
	VERTEX_ARRAY =				0x8074
	NORMAL_ARRAY =				0x8075
	COLOR_ARRAY =				0x8076
	INDEX_ARRAY =				0x8077
	TEXTURE_COORD_ARRAY =			0x8078
	EDGE_FLAG_ARRAY =			0x8079
	VERTEX_ARRAY_SIZE =			0x807A
	VERTEX_ARRAY_TYPE =			0x807B
	VERTEX_ARRAY_STRIDE =			0x807C
	NORMAL_ARRAY_TYPE =			0x807E
	NORMAL_ARRAY_STRIDE =			0x807F
	COLOR_ARRAY_SIZE =			0x8081
	COLOR_ARRAY_TYPE =			0x8082
	COLOR_ARRAY_STRIDE =			0x8083
	INDEX_ARRAY_TYPE =			0x8085
	INDEX_ARRAY_STRIDE =			0x8086
	TEXTURE_COORD_ARRAY_SIZE =		0x8088
	TEXTURE_COORD_ARRAY_TYPE =		0x8089
	TEXTURE_COORD_ARRAY_STRIDE =		0x808A
	EDGE_FLAG_ARRAY_STRIDE =		0x808C
	VERTEX_ARRAY_POINTER =			0x808E
	NORMAL_ARRAY_POINTER =			0x808F
	COLOR_ARRAY_POINTER =			0x8090
	INDEX_ARRAY_POINTER =			0x8091
	TEXTURE_COORD_ARRAY_POINTER =		0x8092
	EDGE_FLAG_ARRAY_POINTER =		0x8093
	V2F =					0x2A20
	V3F =					0x2A21
	C4UB_V2F =				0x2A22
	C4UB_V3F =				0x2A23
	C3F_V3F =				0x2A24
	N3F_V3F =				0x2A25
	C4F_N3F_V3F =				0x2A26
	T2F_V3F =				0x2A27
	T4F_V4F =				0x2A28
	T2F_C4UB_V3F =				0x2A29
	T2F_C3F_V3F =				0x2A2A
	T2F_N3F_V3F =				0x2A2B
	T2F_C4F_N3F_V3F =			0x2A2C
	T4F_C4F_N3F_V4F =			0x2A2D

	/* Matrix Mode */
	MATRIX_MODE =				0x0BA0
	MODELVIEW =				0x1700
	PROJECTION =				0x1701
	TEXTURE =				0x1702

	/* Points */
	POINT_SMOOTH =				0x0B10
	POINT_SIZE =				0x0B11
	POINT_SIZE_GRANULARITY =		0x0B13
	POINT_SIZE_RANGE =			0x0B12

	/* Lines */
	LINE_SMOOTH =				0x0B20
	LINE_STIPPLE =				0x0B24
	LINE_STIPPLE_PATTERN =			0x0B25
	LINE_STIPPLE_REPEAT =			0x0B26
	LINE_WIDTH =				0x0B21
	LINE_WIDTH_GRANULARITY =		0x0B23
	LINE_WIDTH_RANGE =			0x0B22

	/* Polygons */
	POINT =					0x1B00
	LINE =					0x1B01
	FILL =					0x1B02
	CW =					0x0900
	CCW =					0x0901
	FRONT =					0x0404
	BACK =					0x0405
	POLYGON_MODE =				0x0B40
	POLYGON_SMOOTH =			0x0B41
	POLYGON_STIPPLE =			0x0B42
	EDGE_FLAG =				0x0B43
	CULL_FACE =				0x0B44
	CULL_FACE_MODE =			0x0B45
	FRONT_FACE =				0x0B46
	POLYGON_OFFSET_FACTOR =			0x8038
	POLYGON_OFFSET_UNITS =			0x2A00
	POLYGON_OFFSET_POINT =			0x2A01
	POLYGON_OFFSET_LINE =			0x2A02
	POLYGON_OFFSET_FILL =			0x8037

	/* Display Lists */
	COMPILE =				0x1300
	COMPILE_AND_EXECUTE =			0x1301
	LIST_BASE =				0x0B32
	LIST_INDEX =				0x0B33
	LIST_MODE =				0x0B30

	/* Depth buffer */
	NEVER =					0x0200
	LESS =					0x0201
	EQUAL =					0x0202
	LEQUAL =				0x0203
	GREATER =				0x0204
	NOTEQUAL =				0x0205
	GEQUAL =				0x0206
	ALWAYS =				0x0207
	DEPTH_TEST =				0x0B71
	DEPTH_BITS =				0x0D56
	DEPTH_CLEAR_VALUE =			0x0B73
	DEPTH_FUNC =				0x0B74
	DEPTH_RANGE =				0x0B70
	DEPTH_WRITEMASK =			0x0B72
	DEPTH_COMPONENT =			0x1902

	/* Lighting */
	LIGHTING =				0x0B50
	LIGHT0 =				0x4000
	LIGHT1 =				0x4001
	LIGHT2 =				0x4002
	LIGHT3 =				0x4003
	LIGHT4 =				0x4004
	LIGHT5 =				0x4005
	LIGHT6 =				0x4006
	LIGHT7 =				0x4007
	SPOT_EXPONENT =				0x1205
	SPOT_CUTOFF =				0x1206
	CONSTANT_ATTENUATION =			0x1207
	LINEAR_ATTENUATION =			0x1208
	QUADRATIC_ATTENUATION =			0x1209
	AMBIENT =				0x1200
	DIFFUSE =				0x1201
	SPECULAR =				0x1202
	SHININESS =				0x1601
	EMISSION =				0x1600
	POSITION =				0x1203
	SPOT_DIRECTION =			0x1204
	AMBIENT_AND_DIFFUSE =			0x1602
	COLOR_INDEXES =				0x1603
	LIGHT_MODEL_TWO_SIDE =			0x0B52
	LIGHT_MODEL_LOCAL_VIEWER =		0x0B51
	LIGHT_MODEL_AMBIENT =			0x0B53
	FRONT_AND_BACK =			0x0408
	SHADE_MODEL =				0x0B54
	FLAT =					0x1D00
	SMOOTH =				0x1D01
	COLOR_MATERIAL =			0x0B57
	COLOR_MATERIAL_FACE =			0x0B55
	COLOR_MATERIAL_PARAMETER =		0x0B56
	NORMALIZE =				0x0BA1

	/* User clipping planes */
	CLIP_PLANE0 =				0x3000
	CLIP_PLANE1 =				0x3001
	CLIP_PLANE2 =				0x3002
	CLIP_PLANE3 =				0x3003
	CLIP_PLANE4 =				0x3004
	CLIP_PLANE5 =				0x3005

	/* Accumulation buffer */
	ACCUM_RED_BITS =			0x0D58
	ACCUM_GREEN_BITS =			0x0D59
	ACCUM_BLUE_BITS =			0x0D5A
	ACCUM_ALPHA_BITS =			0x0D5B
	ACCUM_CLEAR_VALUE =			0x0B80
	ACCUM =					0x0100
	ADD =					0x0104
	LOAD =					0x0101
	MULT =					0x0103
	RETURN =				0x0102

	/* Alpha testing */
	ALPHA_TEST =				0x0BC0
	ALPHA_TEST_REF =			0x0BC2
	ALPHA_TEST_FUNC =			0x0BC1

	/* Blending */
	BLEND =					0x0BE2
	BLEND_SRC =				0x0BE1
	BLEND_DST =				0x0BE0
	ZERO =					0
	ONE =					1
	SRC_COLOR =				0x0300
	ONE_MINUS_SRC_COLOR =			0x0301
	SRC_ALPHA =				0x0302
	ONE_MINUS_SRC_ALPHA =			0x0303
	DST_ALPHA =				0x0304
	ONE_MINUS_DST_ALPHA =			0x0305
	DST_COLOR =				0x0306
	ONE_MINUS_DST_COLOR =			0x0307
	SRC_ALPHA_SATURATE =			0x0308

	/* Render Mode */
	FEEDBACK =				0x1C01
	RENDER =				0x1C00
	SELECT =				0x1C02

	/* Feedback */
	_2D =					0x0600
	_3D =					0x0601
	_3D_COLOR =				0x0602
	_3D_COLOR_TEXTURE =			0x0603
	_4D_COLOR_TEXTURE =			0x0604
	POINT_TOKEN =				0x0701
	LINE_TOKEN =				0x0702
	LINE_RESET_TOKEN =			0x0707
	POLYGON_TOKEN =				0x0703
	BITMAP_TOKEN =				0x0704
	DRAW_PIXEL_TOKEN =			0x0705
	COPY_PIXEL_TOKEN =			0x0706
	PASS_THROUGH_TOKEN =			0x0700
	FEEDBACK_BUFFER_POINTER =		0x0DF0
	FEEDBACK_BUFFER_SIZE =			0x0DF1
	FEEDBACK_BUFFER_TYPE =			0x0DF2

	/* Selection */
	SELECTION_BUFFER_POINTER =		0x0DF3
	SELECTION_BUFFER_SIZE =			0x0DF4

	/* Fog */
	FOG =					0x0B60
	FOG_MODE =				0x0B65
	FOG_DENSITY =				0x0B62
	FOG_COLOR =				0x0B66
	FOG_INDEX =				0x0B61
	FOG_START =				0x0B63
	FOG_END =				0x0B64
	LINEAR =				0x2601
	EXP =					0x0800
	EXP2 =					0x0801

	/* Logic Ops */
	LOGIC_OP =				0x0BF1
	INDEX_LOGIC_OP =			0x0BF1
	COLOR_LOGIC_OP =			0x0BF2
	LOGIC_OP_MODE =				0x0BF0
	CLEAR =					0x1500
	SET =					0x150F
	COPY =					0x1503
	COPY_INVERTED =				0x150C
	NOOP =					0x1505
	INVERT =				0x150A
	AND =					0x1501
	NAND =					0x150E
	OR =					0x1507
	NOR =					0x1508
	XOR =					0x1506
	EQUIV =					0x1509
	AND_REVERSE =				0x1502
	AND_INVERTED =				0x1504
	OR_REVERSE =				0x150B
	OR_INVERTED =				0x150D

	/* Stencil */
	STENCIL_BITS =				0x0D57
	STENCIL_TEST =				0x0B90
	STENCIL_CLEAR_VALUE =			0x0B91
	STENCIL_FUNC =				0x0B92
	STENCIL_VALUE_MASK =			0x0B93
	STENCIL_FAIL =				0x0B94
	STENCIL_PASS_DEPTH_FAIL =		0x0B95
	STENCIL_PASS_DEPTH_PASS =		0x0B96
	STENCIL_REF =				0x0B97
	STENCIL_WRITEMASK =			0x0B98
	STENCIL_INDEX =				0x1901
	KEEP =					0x1E00
	REPLACE =				0x1E01
	INCR =					0x1E02
	DECR =					0x1E03

	/* Buffers, Pixel Drawing/Reading */
	NONE =					0
	LEFT =					0x0406
	RIGHT =					0x0407
	FRONT_LEFT =				0x0400
	FRONT_RIGHT =				0x0401
	BACK_LEFT =				0x0402
	BACK_RIGHT =				0x0403
	AUX0 =					0x0409
	AUX1 =					0x040A
	AUX2 =					0x040B
	AUX3 =					0x040C
	COLOR_INDEX =				0x1900
	RED =					0x1903
	GREEN =					0x1904
	BLUE =					0x1905
	ALPHA =					0x1906
	LUMINANCE =				0x1909
	LUMINANCE_ALPHA =			0x190A
	ALPHA_BITS =				0x0D55
	RED_BITS =				0x0D52
	GREEN_BITS =				0x0D53
	BLUE_BITS =				0x0D54
	INDEX_BITS =				0x0D51
	SUBPIXEL_BITS =				0x0D50
	AUX_BUFFERS =				0x0C00
	READ_BUFFER =				0x0C02
	DRAW_BUFFER =				0x0C01
	DOUBLEBUFFER =				0x0C32
	STEREO =				0x0C33
	BITMAP =				0x1A00
	COLOR =					0x1800
	DEPTH =					0x1801
	STENCIL =				0x1802
	DITHER =				0x0BD0
	RGB =					0x1907
	RGBA =					0x1908

	/* Implementation limits */
	MAX_LIST_NESTING =			0x0B31
	MAX_EVAL_ORDER =			0x0D30
	MAX_LIGHTS =				0x0D31
	MAX_CLIP_PLANES =			0x0D32
	MAX_TEXTURE_SIZE =			0x0D33
	MAX_PIXEL_MAP_TABLE =			0x0D34
	MAX_ATTRIB_STACK_DEPTH =		0x0D35
	MAX_MODELVIEW_STACK_DEPTH =		0x0D36
	MAX_NAME_STACK_DEPTH =			0x0D37
	MAX_PROJECTION_STACK_DEPTH =		0x0D38
	MAX_TEXTURE_STACK_DEPTH =		0x0D39
	MAX_VIEWPORT_DIMS =			0x0D3A
	MAX_CLIENT_ATTRIB_STACK_DEPTH =		0x0D3B

	/* Gets */
	ATTRIB_STACK_DEPTH =			0x0BB0
	CLIENT_ATTRIB_STACK_DEPTH =		0x0BB1
	COLOR_CLEAR_VALUE =			0x0C22
	COLOR_WRITEMASK =			0x0C23
	CURRENT_INDEX =				0x0B01
	CURRENT_COLOR =				0x0B00
	CURRENT_NORMAL =			0x0B02
	CURRENT_RASTER_COLOR =			0x0B04
	CURRENT_RASTER_DISTANCE =		0x0B09
	CURRENT_RASTER_INDEX =			0x0B05
	CURRENT_RASTER_POSITION =		0x0B07
	CURRENT_RASTER_TEXTURE_COORDS =		0x0B06
	CURRENT_RASTER_POSITION_VALID =		0x0B08
	CURRENT_TEXTURE_COORDS =		0x0B03
	INDEX_CLEAR_VALUE =			0x0C20
	INDEX_MODE =				0x0C30
	INDEX_WRITEMASK =			0x0C21
	MODELVIEW_MATRIX =			0x0BA6
	MODELVIEW_STACK_DEPTH =			0x0BA3
	NAME_STACK_DEPTH =			0x0D70
	PROJECTION_MATRIX =			0x0BA7
	PROJECTION_STACK_DEPTH =		0x0BA4
	RENDER_MODE =				0x0C40
	RGBA_MODE =				0x0C31
	TEXTURE_MATRIX =			0x0BA8
	TEXTURE_STACK_DEPTH =			0x0BA5
	VIEWPORT =				0x0BA2

	/* Evaluators */
	AUTO_NORMAL =				0x0D80
	MAP1_COLOR_4 =				0x0D90
	MAP1_INDEX =				0x0D91
	MAP1_NORMAL =				0x0D92
	MAP1_TEXTURE_COORD_1 =			0x0D93
	MAP1_TEXTURE_COORD_2 =			0x0D94
	MAP1_TEXTURE_COORD_3 =			0x0D95
	MAP1_TEXTURE_COORD_4 =			0x0D96
	MAP1_VERTEX_3 =				0x0D97
	MAP1_VERTEX_4 =				0x0D98
	MAP2_COLOR_4 =				0x0DB0
	MAP2_INDEX =				0x0DB1
	MAP2_NORMAL =				0x0DB2
	MAP2_TEXTURE_COORD_1 =			0x0DB3
	MAP2_TEXTURE_COORD_2 =			0x0DB4
	MAP2_TEXTURE_COORD_3 =			0x0DB5
	MAP2_TEXTURE_COORD_4 =			0x0DB6
	MAP2_VERTEX_3 =				0x0DB7
	MAP2_VERTEX_4 =				0x0DB8
	MAP1_GRID_DOMAIN =			0x0DD0
	MAP1_GRID_SEGMENTS =			0x0DD1
	MAP2_GRID_DOMAIN =			0x0DD2
	MAP2_GRID_SEGMENTS =			0x0DD3
	COEFF =					0x0A00
	ORDER =					0x0A01
	DOMAIN =				0x0A02

	/* Hints */
	PERSPECTIVE_CORRECTION_HINT =		0x0C50
	POINT_SMOOTH_HINT =			0x0C51
	LINE_SMOOTH_HINT =			0x0C52
	POLYGON_SMOOTH_HINT =			0x0C53
	FOG_HINT =				0x0C54
	DONT_CARE =				0x1100
	FASTEST =				0x1101
	NICEST =				0x1102

	/* Scissor box */
	SCISSOR_BOX =				0x0C10
	SCISSOR_TEST =				0x0C11

	/* Pixel Mode / Transfer */
	MAP_COLOR =				0x0D10
	MAP_STENCIL =				0x0D11
	INDEX_SHIFT =				0x0D12
	INDEX_OFFSET =				0x0D13
	RED_SCALE =				0x0D14
	RED_BIAS =				0x0D15
	GREEN_SCALE =				0x0D18
	GREEN_BIAS =				0x0D19
	BLUE_SCALE =				0x0D1A
	BLUE_BIAS =				0x0D1B
	ALPHA_SCALE =				0x0D1C
	ALPHA_BIAS =				0x0D1D
	DEPTH_SCALE =				0x0D1E
	DEPTH_BIAS =				0x0D1F
	PIXEL_MAP_S_TO_S_SIZE =			0x0CB1
	PIXEL_MAP_I_TO_I_SIZE =			0x0CB0
	PIXEL_MAP_I_TO_R_SIZE =			0x0CB2
	PIXEL_MAP_I_TO_G_SIZE =			0x0CB3
	PIXEL_MAP_I_TO_B_SIZE =			0x0CB4
	PIXEL_MAP_I_TO_A_SIZE =			0x0CB5
	PIXEL_MAP_R_TO_R_SIZE =			0x0CB6
	PIXEL_MAP_G_TO_G_SIZE =			0x0CB7
	PIXEL_MAP_B_TO_B_SIZE =			0x0CB8
	PIXEL_MAP_A_TO_A_SIZE =			0x0CB9
	PIXEL_MAP_S_TO_S =			0x0C71
	PIXEL_MAP_I_TO_I =			0x0C70
	PIXEL_MAP_I_TO_R =			0x0C72
	PIXEL_MAP_I_TO_G =			0x0C73
	PIXEL_MAP_I_TO_B =			0x0C74
	PIXEL_MAP_I_TO_A =			0x0C75
	PIXEL_MAP_R_TO_R =			0x0C76
	PIXEL_MAP_G_TO_G =			0x0C77
	PIXEL_MAP_B_TO_B =			0x0C78
	PIXEL_MAP_A_TO_A =			0x0C79
	PACK_ALIGNMENT =			0x0D05
	PACK_LSB_FIRST =			0x0D01
	PACK_ROW_LENGTH =			0x0D02
	PACK_SKIP_PIXELS =			0x0D04
	PACK_SKIP_ROWS =			0x0D03
	PACK_SWAP_BYTES =			0x0D00
	UNPACK_ALIGNMENT =			0x0CF5
	UNPACK_LSB_FIRST =			0x0CF1
	UNPACK_ROW_LENGTH =			0x0CF2
	UNPACK_SKIP_PIXELS =			0x0CF4
	UNPACK_SKIP_ROWS =			0x0CF3
	UNPACK_SWAP_BYTES =			0x0CF0
	ZOOM_X =				0x0D16
	ZOOM_Y =				0x0D17

	/* Texture mapping */
	TEXTURE_ENV =				0x2300
	TEXTURE_ENV_MODE =			0x2200
	TEXTURE_1D =				0x0DE0
	TEXTURE_2D =				0x0DE1
	TEXTURE_WRAP_S =			0x2802
	TEXTURE_WRAP_T =			0x2803
	TEXTURE_MAG_FILTER =			0x2800
	TEXTURE_MIN_FILTER =			0x2801
	TEXTURE_ENV_COLOR =			0x2201
	TEXTURE_GEN_S =				0x0C60
	TEXTURE_GEN_T =				0x0C61
	TEXTURE_GEN_R =				0x0C62
	TEXTURE_GEN_Q =				0x0C63
	TEXTURE_GEN_MODE =			0x2500
	TEXTURE_BORDER_COLOR =			0x1004
	TEXTURE_WIDTH =				0x1000
	TEXTURE_HEIGHT =			0x1001
	TEXTURE_BORDER =			0x1005
	TEXTURE_COMPONENTS =			0x1003
	TEXTURE_RED_SIZE =			0x805C
	TEXTURE_GREEN_SIZE =			0x805D
	TEXTURE_BLUE_SIZE =			0x805E
	TEXTURE_ALPHA_SIZE =			0x805F
	TEXTURE_LUMINANCE_SIZE =		0x8060
	TEXTURE_INTENSITY_SIZE =		0x8061
	NEAREST_MIPMAP_NEAREST =		0x2700
	NEAREST_MIPMAP_LINEAR =			0x2702
	LINEAR_MIPMAP_NEAREST =			0x2701
	LINEAR_MIPMAP_LINEAR =			0x2703
	OBJECT_LINEAR =				0x2401
	OBJECT_PLANE =				0x2501
	EYE_LINEAR =				0x2400
	EYE_PLANE =				0x2502
	SPHERE_MAP =				0x2402
	DECAL =					0x2101
	MODULATE =				0x2100
	NEAREST =				0x2600
	REPEAT =				0x2901
	CLAMP =					0x2900
	S =					0x2000
	T =					0x2001
	R =					0x2002
	Q =					0x2003

	/* Utility */
	VENDOR =				0x1F00
	RENDERER =				0x1F01
	VERSION =				0x1F02
	EXTENSIONS =				0x1F03

	/* Errors */
	NO_ERROR =				0
	INVALID_ENUM =				0x0500
	INVALID_VALUE =				0x0501
	INVALID_OPERATION =			0x0502
	STACK_OVERFLOW =			0x0503
	STACK_UNDERFLOW =			0x0504
	OUT_OF_MEMORY =				0x0505

	/* glPush/PopAttrib bits */
	CURRENT_BIT =				0x00000001
	POINT_BIT =				0x00000002
	LINE_BIT =				0x00000004
	POLYGON_BIT =				0x00000008
	POLYGON_STIPPLE_BIT =			0x00000010
	PIXEL_MODE_BIT =			0x00000020
	LIGHTING_BIT =				0x00000040
	FOG_BIT =				0x00000080
	DEPTH_BUFFER_BIT =			0x00000100
	ACCUM_BUFFER_BIT =			0x00000200
	STENCIL_BUFFER_BIT =			0x00000400
	VIEWPORT_BIT =				0x00000800
	TRANSFORM_BIT =				0x00001000
	ENABLE_BIT =				0x00002000
	COLOR_BUFFER_BIT =			0x00004000
	HINT_BIT =				0x00008000
	EVAL_BIT =				0x00010000
	LIST_BIT =				0x00020000
	TEXTURE_BIT =				0x00040000
	SCISSOR_BIT =				0x00080000
	ALL_ATTRIB_BITS =			0x000FFFFF


	/* OpenGL 1.1 */
	PROXY_TEXTURE_1D =			0x8063
	PROXY_TEXTURE_2D =			0x8064
	TEXTURE_PRIORITY =			0x8066
	TEXTURE_RESIDENT =			0x8067
	TEXTURE_BINDING_1D =			0x8068
	TEXTURE_BINDING_2D =			0x8069
	TEXTURE_INTERNAL_FORMAT =		0x1003
	ALPHA4 =				0x803B
	ALPHA8 =				0x803C
	ALPHA12 =				0x803D
	ALPHA16 =				0x803E
	LUMINANCE4 =				0x803F
	LUMINANCE8 =				0x8040
	LUMINANCE12 =				0x8041
	LUMINANCE16 =				0x8042
	LUMINANCE4_ALPHA4 =			0x8043
	LUMINANCE6_ALPHA2 =			0x8044
	LUMINANCE8_ALPHA8 =			0x8045
	LUMINANCE12_ALPHA4 =			0x8046
	LUMINANCE12_ALPHA12 =			0x8047
	LUMINANCE16_ALPHA16 =			0x8048
	INTENSITY =				0x8049
	INTENSITY4 =				0x804A
	INTENSITY8 =				0x804B
	INTENSITY12 =				0x804C
	INTENSITY16 =				0x804D
	R3_G3_B2 =				0x2A10
	RGB4 =					0x804F
	RGB5 =					0x8050
	RGB8 =					0x8051
	RGB10 =					0x8052
	RGB12 =					0x8053
	RGB16 =					0x8054
	RGBA2 =					0x8055
	RGBA4 =					0x8056
	RGB5_A1 =				0x8057
	RGBA8 =					0x8058
	RGB10_A2 =				0x8059
	RGBA12 =				0x805A
	RGBA16 =				0x805B
	CLIENT_PIXEL_STORE_BIT =		0x00000001
	CLIENT_VERTEX_ARRAY_BIT =		0x00000002
	ALL_CLIENT_ATTRIB_BITS =		0xFFFFFFFF
	CLIENT_ALL_ATTRIB_BITS =		0xFFFFFFFF

	RESCALE_NORMAL =			0x803A
	CLAMP_TO_EDGE =				0x812F
	MAX_ELEMENTS_VERTICES =			0x80E8
	MAX_ELEMENTS_INDICES =			0x80E9
	BGR =					0x80E0
	BGRA =					0x80E1
	UNSIGNED_BYTE_3_3_2 =			0x8032
	UNSIGNED_BYTE_2_3_3_REV =		0x8362
	UNSIGNED_SHORT_5_6_5 =			0x8363
	UNSIGNED_SHORT_5_6_5_REV =		0x8364
	UNSIGNED_SHORT_4_4_4_4 =		0x8033
	UNSIGNED_SHORT_4_4_4_4_REV =		0x8365
	UNSIGNED_SHORT_5_5_5_1 =		0x8034
	UNSIGNED_SHORT_1_5_5_5_REV =		0x8366
	UNSIGNED_INT_8_8_8_8 =			0x8035
	UNSIGNED_INT_8_8_8_8_REV =		0x8367
	UNSIGNED_INT_10_10_10_2 =		0x8036
	UNSIGNED_INT_2_10_10_10_REV =		0x8368
	LIGHT_MODEL_COLOR_CONTROL =		0x81F8
	SINGLE_COLOR =				0x81F9
	SEPARATE_SPECULAR_COLOR =		0x81FA
	TEXTURE_MIN_LOD =			0x813A
	TEXTURE_MAX_LOD =			0x813B
	TEXTURE_BASE_LEVEL =			0x813C
	TEXTURE_MAX_LEVEL =			0x813D
	SMOOTH_POINT_SIZE_RANGE =		0x0B12
	SMOOTH_POINT_SIZE_GRANULARITY =		0x0B13
	SMOOTH_LINE_WIDTH_RANGE =		0x0B22
	SMOOTH_LINE_WIDTH_GRANULARITY =		0x0B23
	ALIASED_POINT_SIZE_RANGE =		0x846D
	ALIASED_LINE_WIDTH_RANGE =		0x846E
	PACK_SKIP_IMAGES =			0x806B
	PACK_IMAGE_HEIGHT =			0x806C
	UNPACK_SKIP_IMAGES =			0x806D
	UNPACK_IMAGE_HEIGHT =			0x806E
	TEXTURE_3D =				0x806F
	PROXY_TEXTURE_3D =			0x8070
	TEXTURE_DEPTH =				0x8071
	TEXTURE_WRAP_R =			0x8072
	MAX_3D_TEXTURE_SIZE =			0x8073
	TEXTURE_BINDING_3D =			0x806A

	CONSTANT_COLOR =			0x8001
	ONE_MINUS_CONSTANT_COLOR =		0x8002
	CONSTANT_ALPHA =			0x8003
	ONE_MINUS_CONSTANT_ALPHA =		0x8004
	COLOR_TABLE =				0x80D0
	POST_CONVOLUTION_COLOR_TABLE =		0x80D1
	POST_COLOR_MATRIX_COLOR_TABLE =		0x80D2
	PROXY_COLOR_TABLE =			0x80D3
	PROXY_POST_CONVOLUTION_COLOR_TABLE =	0x80D4
	PROXY_POST_COLOR_MATRIX_COLOR_TABLE =	0x80D5
	COLOR_TABLE_SCALE =			0x80D6
	COLOR_TABLE_BIAS =			0x80D7
	COLOR_TABLE_FORMAT =			0x80D8
	COLOR_TABLE_WIDTH =			0x80D9
	COLOR_TABLE_RED_SIZE =			0x80DA
	COLOR_TABLE_GREEN_SIZE =		0x80DB
	COLOR_TABLE_BLUE_SIZE =			0x80DC
	COLOR_TABLE_ALPHA_SIZE =		0x80DD
	COLOR_TABLE_LUMINANCE_SIZE =		0x80DE
	COLOR_TABLE_INTENSITY_SIZE =		0x80DF
	CONVOLUTION_1D =			0x8010
	CONVOLUTION_2D =			0x8011
	SEPARABLE_2D =				0x8012
	CONVOLUTION_BORDER_MODE =		0x8013
	CONVOLUTION_FILTER_SCALE =		0x8014
	CONVOLUTION_FILTER_BIAS =		0x8015
	REDUCE =				0x8016
	CONVOLUTION_FORMAT =			0x8017
	CONVOLUTION_WIDTH =			0x8018
	CONVOLUTION_HEIGHT =			0x8019
	MAX_CONVOLUTION_WIDTH =			0x801A
	MAX_CONVOLUTION_HEIGHT =		0x801B
	POST_CONVOLUTION_RED_SCALE =		0x801C
	POST_CONVOLUTION_GREEN_SCALE =		0x801D
	POST_CONVOLUTION_BLUE_SCALE =		0x801E
	POST_CONVOLUTION_ALPHA_SCALE =		0x801F
	POST_CONVOLUTION_RED_BIAS =		0x8020
	POST_CONVOLUTION_GREEN_BIAS =		0x8021
	POST_CONVOLUTION_BLUE_BIAS =		0x8022
	POST_CONVOLUTION_ALPHA_BIAS =		0x8023
	CONSTANT_BORDER =			0x8151
	REPLICATE_BORDER =			0x8153
	CONVOLUTION_BORDER_COLOR =		0x8154
	COLOR_MATRIX =				0x80B1
	COLOR_MATRIX_STACK_DEPTH =		0x80B2
	MAX_COLOR_MATRIX_STACK_DEPTH =		0x80B3
	POST_COLOR_MATRIX_RED_SCALE =		0x80B4
	POST_COLOR_MATRIX_GREEN_SCALE =		0x80B5
	POST_COLOR_MATRIX_BLUE_SCALE =		0x80B6
	POST_COLOR_MATRIX_ALPHA_SCALE =		0x80B7
	POST_COLOR_MATRIX_RED_BIAS =		0x80B8
	POST_COLOR_MATRIX_GREEN_BIAS =		0x80B9
	POST_COLOR_MATRIX_BLUE_BIAS =		0x80BA
	POST_COLOR_MATRIX_ALPHA_BIAS =		0x80BB
	HISTOGRAM =				0x8024
	PROXY_HISTOGRAM =			0x8025
	HISTOGRAM_WIDTH =			0x8026
	HISTOGRAM_FORMAT =			0x8027
	HISTOGRAM_RED_SIZE =			0x8028
	HISTOGRAM_GREEN_SIZE =			0x8029
	HISTOGRAM_BLUE_SIZE =			0x802A
	HISTOGRAM_ALPHA_SIZE =			0x802B
	HISTOGRAM_LUMINANCE_SIZE =		0x802C
	HISTOGRAM_SINK =			0x802D
	MINMAX =				0x802E
	MINMAX_FORMAT =				0x802F
	MINMAX_SINK =				0x8030
	TABLE_TOO_LARGE =			0x8031
	BLEND_EQUATION =			0x8009
	MIN =					0x8007
	MAX =					0x8008
	FUNC_ADD =				0x8006
	FUNC_SUBTRACT =				0x800A
	FUNC_REVERSE_SUBTRACT =			0x800B
	BLEND_COLOR =				0x8005

	/*
	 * OpenGL 1.3
	 */

	/* multitexture */
	TEXTURE0 =				0x84C0
	TEXTURE1 =				0x84C1
	TEXTURE2 =				0x84C2
	TEXTURE3 =				0x84C3
	TEXTURE4 =				0x84C4
	TEXTURE5 =				0x84C5
	TEXTURE6 =				0x84C6
	TEXTURE7 =				0x84C7
	TEXTURE8 =				0x84C8
	TEXTURE9 =				0x84C9
	TEXTURE10 =				0x84CA
	TEXTURE11 =				0x84CB
	TEXTURE12 =				0x84CC
	TEXTURE13 =				0x84CD
	TEXTURE14 =				0x84CE
	TEXTURE15 =				0x84CF
	TEXTURE16 =				0x84D0
	TEXTURE17 =				0x84D1
	TEXTURE18 =				0x84D2
	TEXTURE19 =				0x84D3
	TEXTURE20 =				0x84D4
	TEXTURE21 =				0x84D5
	TEXTURE22 =				0x84D6
	TEXTURE23 =				0x84D7
	TEXTURE24 =				0x84D8
	TEXTURE25 =				0x84D9
	TEXTURE26 =				0x84DA
	TEXTURE27 =				0x84DB
	TEXTURE28 =				0x84DC
	TEXTURE29 =				0x84DD
	TEXTURE30 =				0x84DE
	TEXTURE31 =				0x84DF
	ACTIVE_TEXTURE =			0x84E0
	CLIENT_ACTIVE_TEXTURE =			0x84E1
	MAX_TEXTURE_UNITS =			0x84E2
	/* texture_cube_map */
	NORMAL_MAP =				0x8511
	REFLECTION_MAP =			0x8512
	TEXTURE_CUBE_MAP =			0x8513
	TEXTURE_BINDING_CUBE_MAP =		0x8514
	TEXTURE_CUBE_MAP_POSITIVE_X =		0x8515
	TEXTURE_CUBE_MAP_NEGATIVE_X =		0x8516
	TEXTURE_CUBE_MAP_POSITIVE_Y =		0x8517
	TEXTURE_CUBE_MAP_NEGATIVE_Y =		0x8518
	TEXTURE_CUBE_MAP_POSITIVE_Z =		0x8519
	TEXTURE_CUBE_MAP_NEGATIVE_Z =		0x851A
	PROXY_TEXTURE_CUBE_MAP =		0x851B
	MAX_CUBE_MAP_TEXTURE_SIZE =		0x851C
	/* texture_compression */
	COMPRESSED_ALPHA =			0x84E9
	COMPRESSED_LUMINANCE =			0x84EA
	COMPRESSED_LUMINANCE_ALPHA =		0x84EB
	COMPRESSED_INTENSITY =			0x84EC
	COMPRESSED_RGB =			0x84ED
	COMPRESSED_RGBA =			0x84EE
	TEXTURE_COMPRESSION_HINT =		0x84EF
	TEXTURE_COMPRESSED_IMAGE_SIZE =		0x86A0
	TEXTURE_COMPRESSED =			0x86A1
	NUM_COMPRESSED_TEXTURE_FORMATS =	0x86A2
	COMPRESSED_TEXTURE_FORMATS =		0x86A3
	/* multisample */
	MULTISAMPLE =				0x809D
	SAMPLE_ALPHA_TO_COVERAGE =		0x809E
	SAMPLE_ALPHA_TO_ONE =			0x809F
	SAMPLE_COVERAGE =			0x80A0
	SAMPLE_BUFFERS =			0x80A8
	SAMPLES =				0x80A9
	SAMPLE_COVERAGE_VALUE =			0x80AA
	SAMPLE_COVERAGE_INVERT =		0x80AB
	MULTISAMPLE_BIT =			0x20000000
	/* transpose_matrix */
	TRANSPOSE_MODELVIEW_MATRIX =		0x84E3
	TRANSPOSE_PROJECTION_MATRIX =		0x84E4
	TRANSPOSE_TEXTURE_MATRIX =		0x84E5
	TRANSPOSE_COLOR_MATRIX =		0x84E6
	/* texture_env_combine */
	COMBINE =				0x8570
	COMBINE_RGB =				0x8571
	COMBINE_ALPHA =				0x8572
	SOURCE0_RGB =				0x8580
	SOURCE1_RGB =				0x8581
	SOURCE2_RGB =				0x8582
	SOURCE0_ALPHA =				0x8588
	SOURCE1_ALPHA =				0x8589
	SOURCE2_ALPHA =				0x858A
	OPERAND0_RGB =				0x8590
	OPERAND1_RGB =				0x8591
	OPERAND2_RGB =				0x8592
	OPERAND0_ALPHA =			0x8598
	OPERAND1_ALPHA =			0x8599
	OPERAND2_ALPHA =			0x859A
	RGB_SCALE =				0x8573
	ADD_SIGNED =				0x8574
	INTERPOLATE =				0x8575
	SUBTRACT =				0x84E7
	CONSTANT =				0x8576
	PRIMARY_COLOR =				0x8577
	PREVIOUS =				0x8578
	/* texture_env_dot3 */
	DOT3_RGB =				0x86AE
	DOT3_RGBA =				0x86AF
	/* texture_border_clamp */
	CLAMP_TO_BORDER =			0x812D

	/*
	 * GL_ARB_multitexture (ARB extension 1 and OpenGL 1.2.1)
	 */
	ARB_multitexture =			1

	TEXTURE0_ARB =				0x84C0
	TEXTURE1_ARB =				0x84C1
	TEXTURE2_ARB =				0x84C2
	TEXTURE3_ARB =				0x84C3
	TEXTURE4_ARB =				0x84C4
	TEXTURE5_ARB =				0x84C5
	TEXTURE6_ARB =				0x84C6
	TEXTURE7_ARB =				0x84C7
	TEXTURE8_ARB =				0x84C8
	TEXTURE9_ARB =				0x84C9
	TEXTURE10_ARB =				0x84CA
	TEXTURE11_ARB =				0x84CB
	TEXTURE12_ARB =				0x84CC
	TEXTURE13_ARB =				0x84CD
	TEXTURE14_ARB =				0x84CE
	TEXTURE15_ARB =				0x84CF
	TEXTURE16_ARB =				0x84D0
	TEXTURE17_ARB =				0x84D1
	TEXTURE18_ARB =				0x84D2
	TEXTURE19_ARB =				0x84D3
	TEXTURE20_ARB =				0x84D4
	TEXTURE21_ARB =				0x84D5
	TEXTURE22_ARB =				0x84D6
	TEXTURE23_ARB =				0x84D7
	TEXTURE24_ARB =				0x84D8
	TEXTURE25_ARB =				0x84D9
	TEXTURE26_ARB =				0x84DA
	TEXTURE27_ARB =				0x84DB
	TEXTURE28_ARB =				0x84DC
	TEXTURE29_ARB =				0x84DD
	TEXTURE30_ARB =				0x84DE
	TEXTURE31_ARB =				0x84DF
	ACTIVE_TEXTURE_ARB =			0x84E0
	CLIENT_ACTIVE_TEXTURE_ARB =		0x84E1
	MAX_TEXTURE_UNITS_ARB =			0x84E2
)

/*
 * Miscellaneous
 */

func ClearIndex(c float32) {
	_c := C.GLfloat(c)
	C.glClearIndex(_c)
}

func ClearColor(red, green, blue, alpha float32) {
	_red := C.GLclampf(red)
	_green := C.GLclampf(green)
	_blue := C.GLclampf(blue)
	_alpha := C.GLclampf(alpha)
	C.glClearColor(_red, _green, _blue, _alpha)

}

func Clear(mask uint32) {
	_mask := C.GLbitfield(mask)
	C.glClear(_mask)
}

func IndexMask(mask uint32) {
	_mask := C.GLuint(mask)
	C.glIndexMask(_mask)
}

func ColorMask(red, green, blue, alpha bool) {
	_red := C.GLboolean(Btoi(red))
	_green := C.GLboolean(Btoi(green))
	_blue := C.GLboolean(Btoi(blue))
	_alpha := C.GLboolean(Btoi(alpha))
	C.glColorMask(_red, _green, _blue, _alpha)
}

func AlphaFunc(fn uint32, ref float32) {
	_fn := C.GLenum(fn)
	_ref := C.GLclampf(ref)
	C.glAlphaFunc(_fn, _ref)
}

func BlendFunc(sfactor, dfactor uint32) {
	_sfactor := C.GLenum(sfactor)
	_dfactor := C.GLenum(dfactor)
	C.glBlendFunc(_sfactor, _dfactor)
}

func LogicOp(opcode uint32) {
	_opcode := C.GLenum(opcode)
	C.glLogicOp(_opcode)
}

func CullFace(mode uint32) {
	_mode := C.GLenum(mode)
	C.glCullFace(_mode)
}

func FrontFace(mode uint32) {
	_mode := C.GLenum(mode)
	C.glFrontFace(_mode)
}

func PointSize(size float32) {
	_size := C.GLfloat(size)
	C.glPointSize(_size)
}

func LineWidth(width float32) {
	_width := C.GLfloat(width)
	C.glLineWidth(_width)
}

func LineStipple(factor int32, pattern uint16) {
	_factor := C.GLint(factor)
	_pattern := C.GLushort(pattern)
	C.glLineStipple(_factor, _pattern)
}

func PolygonMode(face, mode uint32) {
	_face := C.GLenum(face)
	_mode := C.GLenum(mode)
	C.glPolygonMode(_face, _mode)
}

func PolygonOffset(factor, units float32) {
	_factor := C.GLfloat(factor)
	_units := C.GLfloat(units)
	C.glPolygonOffset(_factor, _units)
}

func PolygonStipple(mask *uint8) {
	_mask := (*C.GLubyte) (unsafe.Pointer(mask))
	C.glPolygonStipple(_mask)
}

func GetPolygonStipple(mask *uint8) {
	_mask := (*C.GLubyte) (unsafe.Pointer(mask))
	C.glGetPolygonStipple(_mask)
}

func EdgeFlag(flag bool) {
	_flag := C.GLboolean(Btoi(flag))
	C.glEdgeFlag(_flag)
}

func EdgeFlagv(flag *bool) {
	_flag := (*C.GLboolean) (unsafe.Pointer(flag))
	C.glEdgeFlagv(_flag)
}

func Scissor(x, y, width, height int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glScissor(_x, _y, _width, _height)
}

func ClipPlane(plane uint32, equation *float64) {
	_plane := C.GLenum(plane)
	_equation := (*C.GLdouble) (equation)
	C.glClipPlane(_plane, _equation)
}

func GetClipPlane(plane uint32, equation *float64) {
	_plane := C.GLenum(plane)
	_equation := (*C.GLdouble) (equation)
	C.glGetClipPlane(_plane, _equation)
}

func DrawBuffer(mode uint32) {
	_mode := C.GLenum(mode)
	C.glDrawBuffer(_mode)
}

func ReadBuffer(mode uint32) {
	_mode := C.GLenum(mode)
	C.glReadBuffer(_mode)
}

func Enable(cap_ uint32) {
	_cap := C.GLenum(cap_)
	C.glEnable(_cap)
}

func Disable(cap_ uint32) {
	_cap := C.GLenum(cap_)
	C.glDisable(_cap)
}

func IsEnabled(cap_ uint32) bool {
	_cap := C.GLenum(cap_)
	return C.glIsEnabled(_cap) > 0
}


func EnableClientState(cap_ uint32) {
	_cap := C.GLenum(cap_)
	C.glEnableClientState(_cap)
}  /* 1.1 */

func DisableClientState(cap_ uint32) {
	_cap := C.GLenum(cap_)
	C.glDisableClientState(_cap)
}  /* 1.1 */


func GetBooleanv(pname uint32, params *bool) {
	_pname := C.GLenum(pname)
	_params := (*C.GLboolean) (unsafe.Pointer(params))
	C.glGetBooleanv(_pname, _params)
}

func GetDoublev(pname uint32, params *float64) {
	_pname := C.GLenum(pname)
	_params := (*C.GLdouble) (unsafe.Pointer(params))
	C.glGetDoublev(_pname, _params)
}

func GetFloatv(pname uint32, params *float32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(params))
	C.glGetFloatv(_pname, _params)
}

func GetIntegerv(pname uint32, params *int32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(params))
	C.glGetIntegerv(_pname, _params)
}


func PushAttrib(mask uint32) {
	_mask := C.GLbitfield(mask)
	C.glPushAttrib(_mask)
}

func PopAttrib() {
	C.glPopAttrib()
}


func PushClientAttrib(mask uint32) {
	_mask := C.GLbitfield(mask)
	C.glPushClientAttrib(_mask)
}  /* 1.1 */

func PopClientAttrib() {
	C.glPopClientAttrib()
}  /* 1.1 */


func RenderMode(mode uint32) int32 {
	_mode := C.GLenum(mode)
	return int32(C.glRenderMode(_mode))
}

func GetError() uint32 {
	return uint32(C.glGetError())
}

func GetString(name uint32) []byte {
	var b []byte
	_name := C.GLenum(name)
	_str := (*C.char) (unsafe.Pointer(C.glGetString(_name)))
	length := int(C.strlen(_str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(unsafe.Pointer(_str))
	return b
}

func Finish() {
	C.glFinish()
}

func Flush() {
	C.glFlush()
}

func Hint(target, mode uint32) {
	_target := C.GLenum(target)
	_mode := C.GLenum(mode)
	C.glHint(_target, _mode)
}


/*
 * Depth Buffer
 */

func ClearDepth(depth float64) {
	_depth := C.GLclampd(depth)
	C.glClearDepth(_depth)
}

func DepthFunc(func_ uint32) {
	_func := C.GLenum(func_)
	C.glDepthFunc(_func)
}

func DepthMask(flag bool) {
	_flag := C.GLboolean(Btoi(flag))
	C.glDepthMask(_flag)
}

func DepthRange(near_val, far_val float64) {
	_near_val := C.GLclampd(near_val)
	_far_val := C.GLclampd(far_val)
	C.glDepthRange(_near_val, _far_val)
}


/*
 * Accumulation Buffer
 */

func ClearAccum(red, green, blue, alpha float32) {
	_red := C.GLfloat(red)
	_green := C.GLfloat(green)
	_blue := C.GLfloat(blue)
	_alpha := C.GLfloat(alpha)
	C.glClearAccum(_red, _green, _blue, _alpha)
}

func Accum(op uint32, value float32) {
	_op := C.GLenum(op)
	_value := C.GLfloat(value)
	C.glAccum(_op, _value)
}


/*
 * Transformation
 */

func MatrixMode(mode uint32) {
	_mode := C.GLenum(mode)
	C.glMatrixMode(_mode)
}

func Ortho(left, right, bottom, top, near_val, far_val float64) {
	_left := C.GLdouble(left)
	_right := C.GLdouble(right)
	_bottom := C.GLdouble(bottom)
	_top := C.GLdouble(top)
	_near_val := C.GLdouble(near_val)
	_far_val := C.GLdouble(far_val)
	C.glOrtho(_left, _right, _bottom, _top, _near_val, _far_val)
}

func Frustum(left, right, bottom, top, near_val, far_val float64) {
	_left := C.GLdouble(left)
	_right := C.GLdouble(right)
	_bottom := C.GLdouble(bottom)
	_top := C.GLdouble(top)
	_near_val := C.GLdouble(near_val)
	_far_val := C.GLdouble(far_val)
	C.glFrustum(_left, _right, _bottom, _top, _near_val, _far_val)
}

func Viewport(x, y, width, height int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glViewport(_x, _y, _width, _height)
}

func PushMatrix() {
	C.glPushMatrix()
}

func PopMatrix() {
	C.glPopMatrix()
}

func LoadIdentity() {
	C.glLoadIdentity()
}

func LoadMatrixd(m []float64) {
	_m := (*C.GLdouble) (unsafe.Pointer(&m[0]))
	C.glLoadMatrixd(_m)
}

func LoadMatrixf(m []float32) {
	_m := (*C.GLfloat) (unsafe.Pointer(&m[0]))
	C.glLoadMatrixf(_m)
}

func MultMatrixd(m []float64) {
	_m := (*C.GLdouble) (unsafe.Pointer(&m[0]))
	C.glMultMatrixd(_m)
}

func MultMatrixf(m []float32) {
	_m := (*C.GLfloat) (unsafe.Pointer(&m[0]))
	C.glMultMatrixf(_m)
}

func Rotated(angle, x, y, z float64) {
	_angle := C.GLdouble(angle)
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glRotated(_angle, _x, _y, _z)
}

func Rotatef(angle, x, y, z float32) {
	_angle := C.GLfloat(angle)
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glRotatef(_angle, _x, _y, _z)
}

func Scaled(x, y, z float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glScaled(_x, _y, _z)
}

func Scalef(x, y, z float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glScalef(_x, _y, _z)
}

func Translated(x, y, z float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glTranslated(_x, _y, _z)
}

func Translatef(x, y, z float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glTranslatef(_x, _y, _z)
}


/*
 * Display Lists
 */

func IsList(list uint32) bool {
	_list := C.GLuint(list)
	return Itob(int(C.glIsList(_list)))
}

func DeleteLists(list uint32, range_ int32) {
	_list := C.GLuint(list)
	_range := C.GLsizei(range_)
	C.glDeleteLists(_list, _range)
}

func GenLists(range_ int32) uint32 {
	_range := C.GLsizei(range_)
	return uint32(C.glGenLists(_range))
}

func NewList(list, mode uint32) {
	_list := C.GLuint(list)
	_mode := C.GLenum(mode)
	C.glNewList(_list, _mode)
}

func EndList() {
	C.glEndList()
}

func CallList(list uint32) {
	_list := C.GLuint(list)
	C.glCallList(_list)
}

func CallLists(n int32, type_ uint32, lists []uint32) {
	_n := C.GLsizei(n)
	_type := C.GLenum(type_)
	_lists := (unsafe.Pointer(&lists[0]))
	C.glCallLists(_n, _type, _lists)
}

func ListBase(base uint32) {
	_base := C.GLuint(base)
	C.glListBase(_base)
}


/*
 * Drawing Functions
 */

func Begin(mode uint32) {
	_mode := C.GLenum(mode)
	C.glBegin(_mode)
}

func End() {
	C.glEnd()
}


func Vertex2d(x, y float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	C.glVertex2d(_x, _y)
}
func Vertex2f(x, y float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	C.glVertex2f(_x, _y)
}
func Vertex2i(x, y int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	C.glVertex2i(_x, _y)
}
func Vertex2s(x, y int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	C.glVertex2s(_x, _y)
}

func Vertex3d(x, y, z float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glVertex3d(_x, _y, _z)
}
func Vertex3f(x, y, z float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glVertex3f(_x, _y, _z)
}
func Vertex3i(x, y, z int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	C.glVertex3i(_x, _y, _z)
}
func Vertex3s(x, y, z int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	_z := C.GLshort(z)
	C.glVertex3s(_x, _y, _z)
}

func Vertex4d(x, y, z, w float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	_w := C.GLdouble(w)
	C.glVertex4d(_x, _y, _z, _w)
}
func Vertex4f(x, y, z, w float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	_w := C.GLfloat(w)
	C.glVertex4f(_x, _y, _z, _w)
}
func Vertex4i(x, y, z, w int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	_w := C.GLint(w)
	C.glVertex4i(_x, _y, _z, _w)
}
func Vertex4s(x, y, z, w int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	_z := C.GLshort(z)
	_w := C.GLshort(w)
	C.glVertex4s(_x, _y, _z, _w)
}

func Vertex2dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glVertex2dv(_v)
}
func Vertex2fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glVertex2fv(_v)
}
func Vertex2iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glVertex2iv(_v)
}
func Vertex2sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glVertex2sv(_v)
}

func Vertex3dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glVertex3dv(_v)
}
func Vertex3fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glVertex3fv(_v)
}
func Vertex3iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glVertex3iv(_v)
}
func Vertex3sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glVertex3sv(_v)
}

func Vertex4dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glVertex4dv(_v)
}
func Vertex4fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glVertex4fv(_v)
}
func Vertex4iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glVertex4iv(_v)
}
func Vertex4sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glVertex4sv(_v)
}

func Normal3b(nx, ny, nz byte) {
	_nx := C.GLbyte(nx)
	_ny := C.GLbyte(ny)
	_nz := C.GLbyte(nz)
	C.glNormal3b(_nx, _ny, _nz)
}
func Normal3d(nx, ny, nz float64) {
	_nx := C.GLdouble(nx)
	_ny := C.GLdouble(ny)
	_nz := C.GLdouble(nz)
	C.glNormal3d(_nx, _ny, _nz)
}
func Normal3f(nx, ny, nz float32) {
	_nx := C.GLfloat(nx)
	_ny := C.GLfloat(ny)
	_nz := C.GLfloat(nz)
	C.glNormal3f(_nx, _ny, _nz)
}
func Normal3i(nx, ny, nz int32) {
	_nx := C.GLint(nx)
	_ny := C.GLint(ny)
	_nz := C.GLint(nz)
	C.glNormal3i(_nx, _ny, _nz)
}
func Normal3s(nx, ny, nz int16) {
	_nx := C.GLshort(nx)
	_ny := C.GLshort(ny)
	_nz := C.GLshort(nz)
	C.glNormal3s(_nx, _ny, _nz)
}

func Normal3bv(v []byte) {
	_v := (*C.GLbyte) (unsafe.Pointer(&v[0]))
	C.glNormal3bv(_v)
}
func Normal3dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glNormal3dv(_v)
}
func Normal3fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glNormal3fv(_v)
}
func Normal3iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glNormal3iv(_v)
}
func Normal3sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glNormal3sv(_v)
}


func Indexd(c float64) {
	_c := C.GLdouble(c)
	C.glIndexd(_c)
}
func Indexf(c float32) {
	_c := C.GLfloat(c)
	C.glIndexf(_c)
}
func Indexi(c int32) {
	_c := C.GLint(c)
	C.glIndexi(_c)
}
func Indexs(c int16) {
	_c := C.GLshort(c)
	C.glIndexs(_c)
}
func Indexub(c byte) {
	_c := C.GLubyte(c)
	C.glIndexub(_c)
}  // 1.1

func Indexdv(c []float64) {
	_c := (*C.GLdouble) (unsafe.Pointer(&c[0]))
	C.glIndexdv(_c)
}
func Indexfv(c []float32) {
	_c := (*C.GLfloat) (unsafe.Pointer(&c[0]))
	C.glIndexfv(_c)
}
func Indexiv(c []int32) {
	_c := (*C.GLint) (unsafe.Pointer(&c[0]))
	C.glIndexiv(_c)
}
func Indexsv(c []int16) {
	_c := (*C.GLshort) (unsafe.Pointer(&c[0]))
	C.glIndexsv(_c)
}
func Indexubv(c []byte) {
	_c := (*C.GLubyte) (unsafe.Pointer(&c[0]))
	C.glIndexubv(_c)
}  // 1.1

func Color3b(red, green, blue byte) {
	_red := C.GLbyte(red)
	_green := C.GLbyte(green)
	_blue := C.GLbyte(blue)
	C.glColor3b(_red, _green, _blue)
}
func Color3d(red, green, blue float64) {
	_red := C.GLdouble(red)
	_green := C.GLdouble(green)
	_blue := C.GLdouble(blue)
	C.glColor3d(_red, _green, _blue)
}
func Color3f(red, green, blue float32) {
	_red := C.GLfloat(red)
	_green := C.GLfloat(green)
	_blue := C.GLfloat(blue)
	C.glColor3f(_red, _green, _blue)
}

func Color3i(red, green, blue int32) {
	_red := C.GLint(red)
	_green := C.GLint(green)
	_blue := C.GLint(blue)
	C.glColor3i(_red, _green, _blue)
}
func Color3s(red, green, blue int16) {
	_red := C.GLshort(red)
	_green := C.GLshort(green)
	_blue := C.GLshort(blue)
	C.glColor3s(_red, _green, _blue)
}
func Color3ub(red, green, blue byte) {
	_red := C.GLubyte(red)
	_green := C.GLubyte(green)
	_blue := C.GLubyte(blue)
	C.glColor3ub(_red, _green, _blue)
}
func Color3ui(red, green, blue uint32) {
	_red := C.GLuint(red)
	_green := C.GLuint(green)
	_blue := C.GLuint(blue)
	C.glColor3ui(_red, _green, _blue)
}
func Color3us(red, green, blue uint16) {
	_red := C.GLushort(red)
	_green := C.GLushort(green)
	_blue := C.GLushort(blue)
	C.glColor3us(_red, _green, _blue)
}

func Color4b(red, green, blue, alpha byte) {
	_red := C.GLbyte(red)
	_green := C.GLbyte(green)
	_blue := C.GLbyte(blue)
	_alpha := C.GLbyte(alpha)
	C.glColor4b(_red, _green, _blue, _alpha)
}
func Color4d(red, green, blue, alpha float64) {
	_red := C.GLdouble(red)
	_green := C.GLdouble(green)
	_blue := C.GLdouble(blue)
	_alpha := C.GLdouble(alpha)
	C.glColor4d(_red, _green, _blue, _alpha)
}
func Color4f(red, green, blue, alpha float32) {
	_red := C.GLfloat(red)
	_green := C.GLfloat(green)
	_blue := C.GLfloat(blue)
	_alpha := C.GLfloat(alpha)
	C.glColor4f(_red, _green, _blue, _alpha)
}
func Color4i(red, green, blue, alpha int32) {
	_red := C.GLint(red)
	_green := C.GLint(green)
	_blue := C.GLint(blue)
	_alpha := C.GLint(alpha)
	C.glColor4i(_red, _green, _blue, _alpha)
}
func Color4s(red, green, blue, alpha int16) {
	_red := C.GLshort(red)
	_green := C.GLshort(green)
	_blue := C.GLshort(blue)
	_alpha := C.GLshort(alpha)
	C.glColor4s(_red, _green, _blue, _alpha)
}
func Color4ub(red, green, blue, alpha byte) {
	_red := C.GLubyte(red)
	_green := C.GLubyte(green)
	_blue := C.GLubyte(blue)
	_alpha := C.GLubyte(alpha)
	C.glColor4ub(_red, _green, _blue, _alpha)
}
func Color4ui(red, green, blue, alpha uint32) {
	_red := C.GLuint(red)
	_green := C.GLuint(green)
	_blue := C.GLuint(blue)
	_alpha := C.GLuint(alpha)
	C.glColor4ui(_red, _green, _blue, _alpha)
}
func Color4us(red, green, blue, alpha uint16) {
	_red := C.GLushort(red)
	_green := C.GLushort(green)
	_blue := C.GLushort(blue)
	_alpha := C.GLushort(alpha)
	C.glColor4us(_red, _green, _blue, _alpha)
}


func Color3bv(v []byte) {
	_v := (*C.GLbyte) (unsafe.Pointer(&v[0]))
	C.glColor3bv(_v)
}
func Color3dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glColor3dv(_v)
}
func Color3fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glColor3fv(_v)
}
func Color3iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glColor3iv(_v)
}
func Color3sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glColor3sv(_v)
}
func Color3ubv(v []byte) {
	_v := (*C.GLubyte) (unsafe.Pointer(&v[0]))
	C.glColor3ubv(_v)
}
func Color3uiv(v []uint32) {
	_v := (*C.GLuint) (unsafe.Pointer(&v[0]))
	C.glColor3uiv(_v)
}
func Color3usv(v []uint16) {
	_v := (*C.GLushort) (unsafe.Pointer(&v[0]))
	C.glColor3usv(_v)
}

func Color4bv(v []byte) {
	_v := (*C.GLbyte) (unsafe.Pointer(&v[0]))
	C.glColor4bv(_v)
}
func Color4dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glColor4dv(_v)
}
func Color4fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glColor4fv(_v)
}
func Color4iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glColor4iv(_v)
}
func Color4sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glColor4sv(_v)
}
func Color4ubv(v []byte) {
	_v := (*C.GLubyte) (unsafe.Pointer(&v[0]))
	C.glColor4ubv(_v)
}
func Color4uiv(v []uint32) {
	_v := (*C.GLuint) (unsafe.Pointer(&v[0]))
	C.glColor4uiv(_v)
}
func Color4usv(v []uint16) {
	_v := (*C.GLushort) (unsafe.Pointer(&v[0]))
	C.glColor4usv(_v)
}


func TexCoord1d(s float64) {
	_s := C.GLdouble(s)
	C.glTexCoord1d(_s)
}
func TexCoord1f(s float32) {
	_s := C.GLfloat(s)
	C.glTexCoord1f(_s)
}
func TexCoord1i(s int32) {
	_s := C.GLint(s)
	C.glTexCoord1i(_s)
}
func TexCoord1s(s int16) {
	_s := C.GLshort(s)
	C.glTexCoord1s(_s)
}

func TexCoord2d(s, t float64) {
	_s := C.GLdouble(s)
	_t := C.GLdouble(t)
	C.glTexCoord2d(_s, _t)
}
func TexCoord2f(s, t float32) {
	_s := C.GLfloat(s)
	_t := C.GLfloat(t)
	C.glTexCoord2f(_s, _t)
}
func TexCoord2i(s, t int32) {
	_s := C.GLint(s)
	_t := C.GLint(t)
	C.glTexCoord2i(_s, _t)
}
func TexCoord2s(s, t int16) {
	_s := C.GLshort(s)
	_t := C.GLshort(t)
	C.glTexCoord2s(_s, _t)
}

func TexCoord3d(s, t, r float64) {
	_s := C.GLdouble(s)
	_t := C.GLdouble(t)
	_r := C.GLdouble(r)
	C.glTexCoord3d(_s, _t, _r)
}
func TexCoord3f(s, t, r float32) {
	_s := C.GLfloat(s)
	_t := C.GLfloat(t)
	_r := C.GLfloat(r)
	C.glTexCoord3f(_s, _t, _r)
}
func TexCoord3i(s, t, r int32) {
	_s := C.GLint(s)
	_t := C.GLint(t)
	_r := C.GLint(r)
	C.glTexCoord3i(_s, _t, _r)
}

func TexCoord3s(s, t, r int16) {
	_s := C.GLshort(s)
	_t := C.GLshort(t)
	_r := C.GLshort(r)
	C.glTexCoord3s(_s, _t, _r)
}

func TexCoord4d(s, t, r, q float64) {
	_s := C.GLdouble(s)
	_t := C.GLdouble(t)
	_r := C.GLdouble(r)
	_q := C.GLdouble(q)
	C.glTexCoord4d(_s, _t, _r, _q)
}

func TexCoord4f(s, t, r, q float32) {
	_s := C.GLfloat(s)
	_t := C.GLfloat(t)
	_r := C.GLfloat(r)
	_q := C.GLfloat(q)
	C.glTexCoord4f(_s, _t, _r, _q)
}

func TexCoord4i(s, t, r, q int32) {
	_s := C.GLint(s)
	_t := C.GLint(t)
	_r := C.GLint(r)
	_q := C.GLint(q)
	C.glTexCoord4i(_s, _t, _r, _q)
}

func TexCoord4s(s, t, r, q int16) {
	_s := C.GLshort(s)
	_t := C.GLshort(t)
	_r := C.GLshort(r)
	_q := C.GLshort(q)
	C.glTexCoord4s(_s, _t, _r, _q)
}

func TexCoord1dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glTexCoord1dv(_v)
}
func TexCoord1fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glTexCoord1fv(_v)
}
func TexCoord1iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glTexCoord1iv(_v)
}
func TexCoord1sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glTexCoord1sv(_v)
}

func TexCoord2dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glTexCoord2dv(_v)
}
func TexCoord2fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glTexCoord2fv(_v)
}
func TexCoord2iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glTexCoord2iv(_v)
}
func TexCoord2sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glTexCoord2sv(_v)
}

func TexCoord3dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glTexCoord3dv(_v)
}
func TexCoord3fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glTexCoord3fv(_v)
}
func TexCoord3iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glTexCoord3iv(_v)
}
func TexCoord3sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glTexCoord3sv(_v)
}

func TexCoord4dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glTexCoord4dv(_v)
}

func TexCoord4fv(v []float64) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glTexCoord4fv(_v)
}

func TexCoord4iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glTexCoord4iv(_v)
}

func TexCoord4sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glTexCoord4sv(_v)
}

func RasterPos2d(x, y float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	C.glRasterPos2d(_x, _y)
}

func RasterPos2f(x, y float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	C.glRasterPos2f(_x, _y)
}

func RasterPos2i(x, y int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	C.glRasterPos2i(_x, _y)
}

func RasterPos2s(x, y int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	C.glRasterPos2s(_x, _y)
}

func RasterPos3d(x, y, z float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	C.glRasterPos3d(_x, _y, _z)
}
func RasterPos3f(x, y, z float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	C.glRasterPos3f(_x, _y, _z)
}
func RasterPos3i(x, y, z int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	C.glRasterPos3i(_x, _y, _z)
}
func RasterPos3s(x, y, z int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	_z := C.GLshort(z)
	C.glRasterPos3s(_x, _y, _z)
}

func RasterPos4d(x, y, z, w float64) {
	_x := C.GLdouble(x)
	_y := C.GLdouble(y)
	_z := C.GLdouble(z)
	_w := C.GLdouble(w)
	C.glRasterPos4d(_x, _y, _z, _w)
}
func RasterPos4f(x, y, z, w float32) {
	_x := C.GLfloat(x)
	_y := C.GLfloat(y)
	_z := C.GLfloat(z)
	_w := C.GLfloat(w)
	C.glRasterPos4f(_x, _y, _z, _w)
}
func RasterPos4i(x, y, z, w int32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_z := C.GLint(z)
	_w := C.GLint(w)
	C.glRasterPos4i(_x, _y, _z, _w)
}
func RasterPos4s(x, y, z, w int16) {
	_x := C.GLshort(x)
	_y := C.GLshort(y)
	_z := C.GLshort(z)
	_w := C.GLshort(w)
	C.glRasterPos4s(_x, _y, _z, _w)
}

func RasterPos2dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glRasterPos2dv(_v)
}
func RasterPos2fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glRasterPos2fv(_v)
}
func RasterPos2iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glRasterPos2iv(_v)
}
func RasterPos2sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glRasterPos2sv(_v)
}

func RasterPos3dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glRasterPos3dv(_v)
}
func RasterPos3fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glRasterPos3fv(_v)
}
func RasterPos3iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glRasterPos3iv(_v)
}
func RasterPos3sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glRasterPos3sv(_v)
}

func RasterPos4dv(v []float64) {
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glRasterPos4dv(_v)
}
func RasterPos4fv(v []float32) {
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glRasterPos4fv(_v)
}
func RasterPos4iv(v []int32) {
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glRasterPos4iv(_v)
}
func RasterPos4sv(v []int16) {
	_v := (*C.GLshort) (unsafe.Pointer(&v[0]))
	C.glRasterPos4sv(_v)
}

func Rectd(x1, y1, x2, y2 float64) {
	_x1 := C.GLdouble(x1)
	_y1 := C.GLdouble(y1)
	_x2 := C.GLdouble(x2)
	_y2 := C.GLdouble(y2)
	C.glRectd(_x1, _y1, _x2, _y2)
}
func Rectf(x1, y1, x2, y2 float32) {
	_x1 := C.GLfloat(x1)
	_y1 := C.GLfloat(y1)
	_x2 := C.GLfloat(x2)
	_y2 := C.GLfloat(y2)
	C.glRectf(_x1, _y1, _x2, _y2)
}
func Recti(x1, y1, x2, y2 int32) {
	_x1 := C.GLint(x1)
	_y1 := C.GLint(y1)
	_x2 := C.GLint(x2)
	_y2 := C.GLint(y2)
	C.glRecti(_x1, _y1, _x2, _y2)
}
func Rects(x1, y1, x2, y2 int16) {
	_x1 := C.GLshort(x1)
	_y1 := C.GLshort(y1)
	_x2 := C.GLshort(x2)
	_y2 := C.GLshort(y2)
	C.glRects(_x1, _y1, _x2, _y2)
}

func Rectdv(v1 []float64, v2 []float64) {
	_v1 := (*C.GLdouble) (unsafe.Pointer(&v1[0]))
	_v2 := (*C.GLdouble) (unsafe.Pointer(&v2[0]))
	C.glRectdv(_v1, _v2)
}
func Rectfv(v1 []float32, v2 []float32) {
	_v1 := (*C.GLfloat) (unsafe.Pointer(&v1[0]))
	_v2 := (*C.GLfloat) (unsafe.Pointer(&v2[0]))
	C.glRectfv(_v1, _v2)
}
func Rectiv(v1 []int32, v2 []int32) {
	_v1 := (*C.GLint) (unsafe.Pointer(&v1[0]))
	_v2 := (*C.GLint) (unsafe.Pointer(&v2[0]))
	C.glRectiv(_v1, _v2)
}
func Rectsv(v1 []int16, v2 []int16) {
	_v1 := (*C.GLshort) (unsafe.Pointer(&v1[0]))
	_v2 := (*C.GLshort) (unsafe.Pointer(&v2[0]))
	C.glRectsv(_v1, _v2)
}


//  Vertex Arrays  (1.1)

func VertexPointer(size int32, type_ uint32, stride int32, ptr unsafe.Pointer) {
	_size := C.GLint(size)
	_type := C.GLenum(type_)
	_stride := C.GLsizei(stride)
	C.glVertexPointer(_size, _type, _stride, ptr)
}

func NormalPointer(type_ uint32, stride int32, ptr unsafe.Pointer) {
	_type := C.GLenum(type_)
	_stride := C.GLsizei(stride)
	C.glNormalPointer(_type, _stride, ptr)
}

func ColorPointer(size int32, type_ uint32, stride int32, ptr unsafe.Pointer) {
	_size := C.GLint(size)
	_type := C.GLenum(type_)
	_stride := C.GLsizei(stride)
	C.glColorPointer(_size, _type, _stride, ptr)
}

func IndexPointer(type_ uint32, stride int32, ptr unsafe.Pointer) {
	_type := C.GLenum(type_)
	_stride := C.GLsizei(stride)
	C.glIndexPointer(_type, _stride, ptr)
}

func TexCoordPointer(size int32, type_ uint32, stride int32, ptr unsafe.Pointer) {
	_size := C.GLint(size)
	_type := C.GLenum(type_)
	_stride := C.GLsizei(stride)
	C.glTexCoordPointer(_size, _type, _stride, ptr)
}

func EdgeFlagPointer(stride int32, ptr unsafe.Pointer) {
	_stride := C.GLsizei(stride)
	C.glEdgeFlagPointer(_stride, ptr)
}

/*
func GetPointerv(GLenum pname, GLvoid **params) {
}
*/

func ArrayElement(i int32) {
	_i := C.GLint(i)
	C.glArrayElement(_i)
}

func DrawArrays(mode uint32, first int32, count int32) {
	_mode := C.GLenum(mode)
	_first := C.GLint(first)
	_count := C.GLsizei(count)
	C.glDrawArrays(_mode, _first, _count)
}

func DrawElements(mode uint32, count int32, type_ uint32, indices unsafe.Pointer) {
	_mode := C.GLenum(mode)
	_count := C.GLsizei(count)
	_type := C.GLenum(type_)
	C.glDrawElements(_mode, _count, _type, indices)
}

func InterleavedArrays(format uint32, stride int32, pointer unsafe.Pointer) {
	_format := C.GLenum(format)
	_stride := C.GLsizei(stride)
	C.glInterleavedArrays(_format, _stride, pointer)
}

// Lighting

func ShadeModel(mode uint32) {
	_mode := C.GLenum(mode)
	C.glShadeModel(_mode)
}

func Lightf(light, pname uint32, param float32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glLightf(_light, _pname, _param)
}

func Lighti(light, pname uint32, param int32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glLighti(_light, _pname, _param)
}

func Lightfv(light, pname uint32, param []float32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := (*C.GLfloat) (unsafe.Pointer(&param[0]))
	C.glLightfv(_light, _pname, _param)
}
func Lightiv(light, pname uint32, param []int32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := (*C.GLint) (unsafe.Pointer(&param[0]))
	C.glLightiv(_light, _pname, _param)
}

func GetLightfv(light, pname uint32, param []float32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := (*C.GLfloat) (unsafe.Pointer(&param[0]))
	C.glGetLightfv(_light, _pname, _param)
}
func GetLightiv(light, pname uint32, param []int32) {
	_light := C.GLenum(light)
	_pname := C.GLenum(pname)
	_param := (*C.GLint) (unsafe.Pointer(&param[0]))
	C.glGetLightiv(_light, _pname, _param)
}

func LightModelf(pname uint32, param float32) {
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glLightModelf(_pname, _param)
}
func LightModeli(pname uint32, param int32) {
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glLightModeli(_pname, _param)
}
func LightModelfv(pname uint32, params []float32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glLightModelfv(_pname, _params)
}
func LightModeliv(pname uint32, params []int32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glLightModeliv(_pname, _params)
}

func Materialf(face, pname uint32, param float32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glMaterialf(_face, _pname, _param)
}
func Materiali(face, pname uint32, param int32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glMateriali(_face, _pname, _param)
}
func Materialfv(face, pname uint32, params []float32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glMaterialfv(_face, _pname, _params)
}
func Materialiv(face, pname uint32, params []int32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glMaterialiv(_face, _pname, _params)
}

func GetMaterialfv(face, pname uint32, params []float32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glGetMaterialfv(_face, _pname, _params)
}

func GetMaterialiv(face, pname uint32, params []int32) {
	_face := C.GLenum(face)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glGetMaterialiv(_face, _pname, _params)
}

func ColorMaterial(face, mode uint32) {
	_face := C.GLenum(face)
	_mode := C.GLenum(mode)
	C.glColorMaterial(_face, _mode)
}

// Raster functions
func PixelZoom(xfactor, yfactor float32) {
	_xfactor := C.GLfloat(xfactor)
	_yfactor := C.GLfloat(yfactor)
	C.glPixelZoom(_xfactor, _yfactor)
}

func PixelStoref(pname uint32, param float32) {
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glPixelStoref(_pname, _param)
}
func PixelStorei(pname uint32, param int32) {
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glPixelStorei(_pname, _param)
}

func PixelTransferf(pname uint32, param float32) {
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glPixelTransferf(_pname, _param)
}
func PixelTransferi(pname uint32, param int32) {
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glPixelTransferi(_pname, _param)
}

func PixelMapfv(map_ uint32, mapsize int32, values []float32) {
	_map := C.GLenum(map_)
	_mapsize := C.GLsizei(mapsize)
	_values := (*C.GLfloat) (unsafe.Pointer(&values[0]))
	C.glPixelMapfv(_map, _mapsize, _values)
}
func PixelMapuiv(map_ uint32, mapsize int32, values []uint32) {
	_map := C.GLenum(map_)
	_mapsize := C.GLsizei(mapsize)
	_values := (*C.GLuint) (unsafe.Pointer(&values[0]))
	C.glPixelMapuiv(_map, _mapsize, _values)
}
func PixelMapusv(map_ uint32, mapsize int32, values []uint16) {
	_map := C.GLenum(map_)
	_mapsize := C.GLsizei(mapsize)
	_values := (*C.GLushort) (unsafe.Pointer(&values[0]))
	C.glPixelMapusv(_map, _mapsize, _values)
}

func GetPixelMapfv(map_ uint32, values []float32) {
	_map := C.GLenum(map_)
	_values := (*C.GLfloat) (unsafe.Pointer(&values[0]))
	C.glGetPixelMapfv(_map, _values)
}
func GetPixelMapuiv(map_ uint32, values []uint32) {
	_map := C.GLenum(map_)
	_values := (*C.GLuint) (unsafe.Pointer(&values[0]))
	C.glGetPixelMapuiv(_map, _values)
}
func GetPixelMapusv(map_ uint32, values []uint16) {
	_map := C.GLenum(map_)
	_values := (*C.GLushort) (unsafe.Pointer(&values[0]))
	C.glGetPixelMapusv(_map, _values)
}

func Bitmap(width, height int32, xorig, yorig, xmove, ymove float32, bitmap []byte) {
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_xorig := C.GLfloat(xorig)
	_yorig := C.GLfloat(yorig)
	_xmove := C.GLfloat(xmove)
	_ymove := C.GLfloat(ymove)
	_bitmap := (*C.GLubyte) (unsafe.Pointer(&bitmap[0]))
	C.glBitmap(_width, _height, _xorig, _yorig, _xmove, _ymove, _bitmap)
}

func ReadPixels(x, y, width, height int32, format, type_ uint32, pixels unsafe.Pointer) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glReadPixels(_x, _y, _width, _height, _format, _type, pixels)
}

func DrawPixels(width, height int32, format, type_ uint32, pixels unsafe.Pointer) {
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glDrawPixels(_width, _height, _format, _type, pixels)
}

func CopyPixels(x, y, width, height int32, type_ uint32) {
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_type := C.GLenum(type_)
	C.glCopyPixels(_x, _y, _width, _height, _type)
}

// Stenciling

func StencilFunc(func_ uint32, ref int32, mask uint32) {
	_func := C.GLenum(func_)
	_ref := C.GLint(ref)
	_mask := C.GLuint(mask)
	C.glStencilFunc(_func, _ref, _mask)
}

func StencilMask(mask uint32) {
	_mask := C.GLuint(mask)
	C.glStencilMask(_mask)
}

func StencilOp(fail, zfail, zpass uint32) {
	_fail := C.GLenum(fail)
	_zfail := C.GLenum(zfail)
	_zpass := C.GLenum(zpass)
	C.glStencilOp(_fail, _zfail, _zpass)
}

func ClearStencil(s int32) {
	_s := C.GLint(s)
	C.glClearStencil(_s)
}



// Texture mapping

func TexGend(coord, pname uint32, param float64) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_param := C.GLdouble(param)
	C.glTexGend(_coord, _pname, _param)
}
func TexGenf(coord, pname uint32, param float32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glTexGenf(_coord, _pname, _param)
}
func TexGeni(coord, pname uint32, param int32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glTexGeni(_coord, _pname, _param)
}

func TexGendv(coord, pname uint32, params []float64) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLdouble) (unsafe.Pointer(&params[0]))
	C.glTexGendv(_coord, _pname, _params)
}
func TexGenfv(coord, pname uint32, params []float32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glTexGenfv(_coord, _pname, _params)
}
func TexGeniv(coord, pname uint32, params []int32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glTexGeniv(_coord, _pname, _params)
}

func GetTexGendv(coord, pname uint32, params []float64) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLdouble) (unsafe.Pointer(&params[0]))
	C.glGetTexGendv(_coord, _pname, _params)
}
func GetTexGenfv(coord, pname uint32, params []float32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glGetTexGenfv(_coord, _pname, _params)
}
func GetTexGeniv(coord, pname uint32, params []int32) {
	_coord := C.GLenum(coord)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glGetTexGeniv(_coord, _pname, _params)
}

func TexEnvf(target, pname uint32, param float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glTexEnvf(_target, _pname, _param)
}
func TexEnvi(target, pname uint32, param int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glTexEnvi(_target, _pname, _param)
}

func TexEnvfv(target, pname uint32, params []float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glTexEnvfv(_target, _pname, _params)
}
func TexEnviv(target, pname uint32, params []int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glTexEnviv(_target, _pname, _params)
}

func GetTexEnvfv(target, pname uint32, params []float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glGetTexEnvfv(_target, _pname, _params)
}
func GetTexEnviv(target, pname uint32, params []int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glGetTexEnviv(_target, _pname, _params)
}

func TexParameterf(target, pname uint32, param float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glTexParameterf(_target, _pname, _param)
}
func TexParameteri(target, pname uint32, param int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glTexParameteri(_target, _pname, _param)
}

func TexParameterfv(target, pname uint32, params []float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glTexParameterfv(_target, _pname, _params)
}
func TexParameteriv(target, pname uint32, params []int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glTexParameteriv(_target, _pname, _params)
}

func GetTexParameterfv(target, pname uint32, params []float32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glGetTexParameterfv(_target, _pname, _params)
}
func GetTexParameteriv(target, pname uint32, params []int32) {
	_target := C.GLenum(target)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glGetTexParameteriv(_target, _pname, _params)
}

func GetTexLevelParameterfv(target uint32, level int32, pname uint32, params []float32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glGetTexLevelParameterfv(_target, _level, _pname, _params)
}
func GetTexLevelParameteriv(target uint32, level int32, pname uint32, params []int32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glGetTexLevelParameteriv(_target, _level, _pname, _params)
}


func TexImage1D(target uint32, level, internalFormat, width, border int32, format, type_ uint32, pixels unsafe.Pointer) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalFormat := C.GLint(internalFormat)
	_width := C.GLsizei(width)
	_border := C.GLint(border)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glTexImage1D(_target, _level, _internalFormat, _width, _border, _format, _type, pixels)
}

func TexImage2D(target uint32, level, internalFormat, width, height, border int32, format, type_ uint32, pixels unsafe.Pointer) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalFormat := C.GLint(internalFormat)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_border := C.GLint(border)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glTexImage2D(_target, _level, _internalFormat, _width, _height, _border, _format, _type, pixels)
}

func GetTexImage(target uint32, level int32, format, type_ uint32, pixels unsafe.Pointer) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glGetTexImage(_target, _level, _format, _type, pixels)
}


// 1.1 functions

func GenTextures(n int32, textures []uint32) {
	_n := C.GLsizei(n)
	_textures := (*C.GLuint) (unsafe.Pointer(&textures[0]))
	C.glGenTextures(_n, _textures)
}

func DeleteTextures(n int32, textures []uint32) {
	_n := C.GLsizei(n)
	_textures := (*C.GLuint) (unsafe.Pointer(&textures[0]))
	C.glDeleteTextures(_n, _textures)
}

func BindTexture(target, texture uint32) {
	_target := C.GLenum(target)
	_texture := C.GLuint(texture)
	C.glBindTexture(_target, _texture)
}

func PrioritizeTextures(n int32, textures []uint32, priorities []float32) {
	_n := C.GLsizei(n)
	_textures := (*C.GLuint) (unsafe.Pointer(&textures[0]))
	_priorities := (*C.GLclampf) (unsafe.Pointer(&priorities[0]))
	C.glPrioritizeTextures(_n, _textures, _priorities)
}

func AreTexturesResident(n int32, textures []uint32, residences []bool) bool {
	_n := C.GLsizei(n)
	_textures := (*C.GLuint) (unsafe.Pointer(&textures[0]))
	_residences := (*C.GLboolean) (unsafe.Pointer(&residences[0]))
	return Itob(int(C.glAreTexturesResident(_n, _textures, _residences)))
}

func IsTexture(texture uint32) bool {
	_texture := C.GLuint(texture)
	return Itob(int(C.glIsTexture(_texture)))
}


func TexSubImage1D(target uint32, level, xoffset, width int32, format, type_ uint32, pixels unsafe.Pointer) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_width := C.GLsizei(width)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glTexSubImage1D(_target, _level, _xoffset, _width, _format, _type, pixels)
}

func TexSubImage2D(target uint32, level, xoffset, yoffset, width, height int32, format, type_ uint32, pixels unsafe.Pointer) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_format := C.GLenum(format)
	_type := C.GLenum(type_)
	C.glTexSubImage2D(_target, _level, _xoffset, _yoffset, _width, _height, _format, _type, pixels)
}

func CopyTexImage1D(target uint32, level int32, internalFormat uint32, x, y, width, border int32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalFormat := C.GLenum(internalFormat)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_border := C.GLint(border)
	C.glCopyTexImage1D(_target, _level, _internalFormat, _x, _y, _width, _border)
}

func CopyTexImage2D(target uint32, level int32, internalFormat uint32, x, y, width, height, border int32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_internalFormat := C.GLenum(internalFormat)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	_border := C.GLint(border)
	C.glCopyTexImage2D(_target, _level, _internalFormat, _x, _y, _width, _height, _border)
}

func CopyTexSubImage1D(target uint32, level, xoffset, x, y, width int32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	C.glCopyTexSubImage1D(_target, _level, _xoffset, _x, _y, _width)
}

func CopyTexSubImage2D(target uint32, level, xoffset, yoffset, x, y, width, height int32) {
	_target := C.GLenum(target)
	_level := C.GLint(level)
	_xoffset := C.GLint(xoffset)
	_yoffset := C.GLint(yoffset)
	_x := C.GLint(x)
	_y := C.GLint(y)
	_width := C.GLsizei(width)
	_height := C.GLsizei(height)
	C.glCopyTexSubImage2D(_target, _level, _xoffset, _yoffset, _x, _y, _width, _height)
}


// Evaluators

func Map1d(target uint32, u1, u2 float64, stride, order int32, points []float64) {
	_target := C.GLenum(target)
	_u1 := C.GLdouble(u1)
	_u2 := C.GLdouble(u2)
	_stride := C.GLint(stride)
	_order := C.GLint(order)
	_points := (*C.GLdouble) (unsafe.Pointer(&points[0]))
	C.glMap1d(_target, _u1, _u2, _stride, _order, _points)
}

func Map1f(target uint32, u1, u2 float32, stride, order int32, points []float32) {
	_target := C.GLenum(target)
	_u1 := C.GLfloat(u1)
	_u2 := C.GLfloat(u2)
	_stride := C.GLint(stride)
	_order := C.GLint(order)
	_points := (*C.GLfloat) (unsafe.Pointer(&points[0]))
	C.glMap1f(_target, _u1, _u2, _stride, _order, _points)
}

func Map2d(target uint32, u1, u2 float64, ustride, uorder int32,
		v1, v2 float64, vstride, vorder int32, points []float64) {
	_target := C.GLenum(target)
	_u1 := C.GLdouble(u1)
	_u2 := C.GLdouble(u2)
	_ustride := C.GLint(ustride)
	_uorder := C.GLint(uorder)
	_v1 := C.GLdouble(v1)
	_v2 := C.GLdouble(v2)
	_vstride := C.GLint(vstride)
	_vorder := C.GLint(vorder)
	_points := (*C.GLdouble) (unsafe.Pointer(&points[0]))
	C.glMap2d(_target, _u1, _u2, _ustride, _uorder, _v1, _v2, _vstride, _vorder, _points)
}

func Map2f(target uint32, u1, u2 float32, ustride, uorder int32,
		v1, v2 float32, vstride, vorder int32, points []float32) {
	_target := C.GLenum(target)
	_u1 := C.GLfloat(u1)
	_u2 := C.GLfloat(u2)
	_ustride := C.GLint(ustride)
	_uorder := C.GLint(uorder)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	_vstride := C.GLint(vstride)
	_vorder := C.GLint(vorder)
	_points := (*C.GLfloat) (unsafe.Pointer(&points[0]))
	C.glMap2f(_target, _u1, _u2, _ustride, _uorder, _v1, _v2, _vstride, _vorder, _points)
}

func GetMapdv(target, query uint32, v []float64) {
	_target := C.GLenum(target)
	_query := C.GLenum(query)
	_v := (*C.GLdouble) (unsafe.Pointer(&v[0]))
	C.glGetMapdv(_target, _query, _v)
}

func GetMapfv(target, query uint32, v []float32) {
	_target := C.GLenum(target)
	_query := C.GLenum(query)
	_v := (*C.GLfloat) (unsafe.Pointer(&v[0]))
	C.glGetMapfv(_target, _query, _v)
}

func GetMapiv(target, query uint32, v []int32) {
	_target := C.GLenum(target)
	_query := C.GLenum(query)
	_v := (*C.GLint) (unsafe.Pointer(&v[0]))
	C.glGetMapiv(_target, _query, _v)
}

func EvalCoord1d(u float64) {
	_u := C.GLdouble(u)
	C.glEvalCoord1d(_u)
}
func EvalCoord1f(u float32) {
	_u := C.GLfloat(u)
	C.glEvalCoord1f(_u)
}

func EvalCoord1dv(u []float64) {
	_u := (*C.GLdouble) (unsafe.Pointer(&u[0]))
	C.glEvalCoord1dv(_u)
}
func EvalCoord1fv(u []float32) {
	_u := (*C.GLfloat) (unsafe.Pointer(&u[0]))
	C.glEvalCoord1fv(_u)
}

func EvalCoord2d(u, v float64) {
	_u := C.GLdouble(u)
	_v := C.GLdouble(v)
	C.glEvalCoord2d(_u, _v)
}
func EvalCoord2f(u, v float32) {
	_u := C.GLfloat(u)
	_v := C.GLfloat(v)
	C.glEvalCoord2f(_u, _v)
}

func EvalCoord2dv(u []float64) {
	_u := (*C.GLdouble) (unsafe.Pointer(&u[0]))
	C.glEvalCoord2dv(_u)
}
func EvalCoord2fv(u []float32) {
	_u := (*C.GLfloat) (unsafe.Pointer(&u[0]))
	C.glEvalCoord2fv(_u)
}

func MapGrid1d(un int32, u1, u2 float64) {
	_un := C.GLint(un)
	_u1 := C.GLdouble(u1)
	_u2 := C.GLdouble(u2)
	C.glMapGrid1d(_un, _u1, _u2)
}
func MapGrid1f(un int32, u1, u2 float32) {
	_un := C.GLint(un)
	_u1 := C.GLfloat(u1)
	_u2 := C.GLfloat(u2)
	C.glMapGrid1f(_un, _u1, _u2)
}

func MapGrid2d(un int64, u1, u2 float64, vn int64, v1, v2 float64) {
	_un := C.GLint(un)
	_u1 := C.GLdouble(u1)
	_u2 := C.GLdouble(u2)
	_vn := C.GLint(vn)
	_v1 := C.GLdouble(v1)
	_v2 := C.GLdouble(v2)
	C.glMapGrid2d(_un, _u1, _u2, _vn, _v1, _v2)
}
func MapGrid2f(un int32, u1, u2 float32, vn int32, v1, v2 float32) {
	_un := C.GLint(un)
	_u1 := C.GLfloat(u1)
	_u2 := C.GLfloat(u2)
	_vn := C.GLint(vn)
	_v1 := C.GLfloat(v1)
	_v2 := C.GLfloat(v2)
	C.glMapGrid2f(_un, _u1, _u2, _vn, _v1, _v2)
}

func EvalPoint1(i int32) {
	_i := C.GLint(i)
	C.glEvalPoint1(_i)
}

func EvalPoint2(i, j int32) {
	_i := C.GLint(i)
	_j := C.GLint(j)
	C.glEvalPoint2(_i, _j)
}

func EvalMesh1(mode uint32, i1, i2 int32) {
	_mode := C.GLenum(mode)
	_i1 := C.GLint(i1)
	_i2 := C.GLint(i2)
	C.glEvalMesh1(_mode, _i1, _i2)
}

func EvalMesh2(mode uint32, i1, i2 int32, j1, j2 int32) {
	_mode := C.GLenum(mode)
	_i1 := C.GLint(i1)
	_i2 := C.GLint(i2)
	_j1 := C.GLint(j1)
	_j2 := C.GLint(j2)
	C.glEvalMesh2(_mode, _i1, _i2, _j1, _j2)
}

// Fog
func Fogf(pname uint32, param float32) {
	_pname := C.GLenum(pname)
	_param := C.GLfloat(param)
	C.glFogf(_pname, _param)
}

func Fogi(pname uint32, param float32) {
	_pname := C.GLenum(pname)
	_param := C.GLint(param)
	C.glFogi(_pname, _param)
}

func Fogfv(pname uint32, params []float32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLfloat) (unsafe.Pointer(&params[0]))
	C.glFogfv(_pname, _params)
}

func Fogiv(pname uint32, params []int32) {
	_pname := C.GLenum(pname)
	_params := (*C.GLint) (unsafe.Pointer(&params[0]))
	C.glFogiv(_pname, _params)
}

// Selection and Feedback
func FeedbackBuffer(size int32, type_ uint32, buffer []float32) {
	_size := C.GLsizei(size)
	_type := C.GLenum(type_)
	_buffer := (*C.GLfloat) (unsafe.Pointer(&buffer[0]))
	C.glFeedbackBuffer(_size, _type,_buffer)
}

func PassThrough(token float32) {
	_token := C.GLfloat(token)
	C.glPassThrough(_token)
}

func SelectBuffer(size int32, buffer []uint32) {
	_size := C.GLsizei(size)
	_buffer := (*C.GLuint) (unsafe.Pointer(&buffer[0]))
	C.glSelectBuffer(_size, _buffer)
}

func InitNames() {
	C.glInitNames()
}

func LoadName(name uint32) {
	_name := C.GLuint(name)
	C.glLoadName(_name)
}

func PushName(name uint32) {
	_name := C.GLuint(name)
	C.glPushName(_name)
}

func PopName() {
	C.glPopName()
}



// OpenGL 1.2

func DrawRangeElements(mode, start, end uint32, count int32, type_ uint32, indices unsafe.Pointer) {
	_mode := C.GLenum(mode)
	_start := C.GLuint(start)
	_end := C.GLuint(end)
	_count := C.GLsizei(count)
	_type := C.GLenum(type_)
	C.glDrawRangeElements(_mode, _start, _end, _count, _type, indices)
}

/*
TexImage3D(GLenum target, GLint level, GLint internalFormat,
	GLsizei width, GLsizei height, GLsizei depth, GLint border,
	GLenum format, GLenum type, const GLvoid *pixels) {
}

TexSubImage3D(GLenum target, GLint level,
	GLint xoffset, GLint yoffset,
	GLint zoffset, GLsizei width,
	GLsizei height, GLsizei depth,
	GLenum format,
	GLenum type, const GLvoid *pixels) {
}

CopyTexSubImage3D( GLenum target, GLint level,
	GLint xoffset, GLint yoffset,
	GLint zoffset, GLint x,
	GLint y, GLsizei width,
	GLsizei height ) {
}

// GL_ARB_imaging

ColorTable( GLenum target, GLenum internalformat, GLsizei width, GLenum format, GLenum type, const GLvoid *table ) {
}

ColorSubTable( GLenum target, GLsizei start, GLsizei count, GLenum format, GLenum type, const GLvoid *data ) {
}

ColorTableParameteriv(GLenum target, GLenum pname, const GLint *params) {
}

ColorTableParameterfv(GLenum target, GLenum pname, const GLfloat *params) {
}

CopyColorSubTable( GLenum target, GLsizei start, GLint x, GLint y, GLsizei width ) {
}

CopyColorTable( GLenum target, GLenum internalformat, GLint x, GLint y, GLsizei width ) {
}

GetColorTable( GLenum target, GLenum format, GLenum type, GLvoid *table ) {
}

GetColorTableParameterfv( GLenum target, GLenum pname,
GLfloat *params ) {
}

GetColorTableParameteriv( GLenum target, GLenum pname, GLint *params ) {
}

BlendEquation( GLenum mode ) {
}

BlendColor( GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha ) {
}

Histogram( GLenum target, GLsizei width, GLenum internalformat, GLboolean sink ) {
}

ResetHistogram( GLenum target ) {
}

GetHistogram( GLenum target, GLboolean reset, GLenum format, GLenum type, GLvoid *values ) {
}

GetHistogramParameterfv( GLenum target, GLenum pname,
						 GLfloat *params ) {
}

GetHistogramParameteriv( GLenum target, GLenum pname,
						 GLint *params ) {
}

Minmax( GLenum target, GLenum internalformat,
				GLboolean sink ) {
}

ResetMinmax( GLenum target ) {
}

GetMinmax( GLenum target, GLboolean reset,
                                   GLenum format, GLenum types,
                                   GLvoid *values ) {
}

GetMinmaxParameterfv( GLenum target, GLenum pname,
					      GLfloat *params ) {
}

GetMinmaxParameteriv( GLenum target, GLenum pname,
					      GLint *params ) {
}

ConvolutionFilter1D( GLenum target,
	GLenum internalformat, GLsizei width, GLenum format, GLenum type,
	const GLvoid *image ) {
}

ConvolutionFilter2D( GLenum target,
	GLenum internalformat, GLsizei width, GLsizei height, GLenum format,
	GLenum type, const GLvoid *image ) {
}

ConvolutionParameterf( GLenum target, GLenum pname,
	GLfloat params ) {
}

ConvolutionParameterfv( GLenum target, GLenum pname,
	const GLfloat *params ) {
}

ConvolutionParameteri( GLenum target, GLenum pname,
	GLint params ) {
}

ConvolutionParameteriv( GLenum target, GLenum pname,
	const GLint *params ) {
}

CopyConvolutionFilter1D( GLenum target,
	GLenum internalformat, GLint x, GLint y, GLsizei width ) {
}

CopyConvolutionFilter2D( GLenum target,
	GLenum internalformat, GLint x, GLint y, GLsizei width,
	GLsizei height) {
}

GetConvolutionFilter( GLenum target, GLenum format,
	GLenum type, GLvoid *image ) {
}

GetConvolutionParameterfv( GLenum target, GLenum pname,
	GLfloat *params ) {
}

GetConvolutionParameteriv( GLenum target, GLenum pname,
	GLint *params ) {
}

SeparableFilter2D( GLenum target,
	GLenum internalformat, GLsizei width, GLsizei height, GLenum format,
	GLenum type, const GLvoid *row, const GLvoid *column ) {
}

GetSeparableFilter( GLenum target, GLenum format,
	GLenum type, GLvoid *row, GLvoid *column, GLvoid *span ) {
}


ActiveTexture( GLenum texture ) {
}

ClientActiveTexture( GLenum texture ) {
}

CompressedTexImage1D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLint border, GLsizei imageSize, const GLvoid *data ) {
}

CompressedTexImage2D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid *data ) {
}

CompressedTexImage3D( GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid *data ) {
}

CompressedTexSubImage1D( GLenum target, GLint level, GLint xoffset, GLsizei width, GLenum format, GLsizei imageSize, const GLvoid *data ) {
}

CompressedTexSubImage2D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid *data ) {
}

CompressedTexSubImage3D( GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid *data ) {
}

GetCompressedTexImage( GLenum target, GLint lod, GLvoid *img ) {
}

MultiTexCoord1d( GLenum target, GLdouble s ) {
}

MultiTexCoord1dv( GLenum target, const GLdouble *v ) {
}

MultiTexCoord1f( GLenum target, GLfloat s ) {
}

MultiTexCoord1fv( GLenum target, const GLfloat *v ) {
}

MultiTexCoord1i( GLenum target, GLint s ) {
}

MultiTexCoord1iv( GLenum target, const GLint *v ) {
}

MultiTexCoord1s( GLenum target, GLshort s ) {
}

MultiTexCoord1sv( GLenum target, const GLshort *v ) {
}

MultiTexCoord2d( GLenum target, GLdouble s, GLdouble t ) {
}

MultiTexCoord2dv( GLenum target, const GLdouble *v ) {
}

MultiTexCoord2f( GLenum target, GLfloat s, GLfloat t ) {
}

MultiTexCoord2fv( GLenum target, const GLfloat *v ) {
}

MultiTexCoord2i( GLenum target, GLint s, GLint t ) {
}

MultiTexCoord2iv( GLenum target, const GLint *v ) {
}

MultiTexCoord2s( GLenum target, GLshort s, GLshort t ) {
}

MultiTexCoord2sv( GLenum target, const GLshort *v ) {
}

MultiTexCoord3d( GLenum target, GLdouble s, GLdouble t, GLdouble r ) {
}

MultiTexCoord3dv( GLenum target, const GLdouble *v ) {
}

MultiTexCoord3f( GLenum target, GLfloat s, GLfloat t, GLfloat r ) {
}

MultiTexCoord3fv( GLenum target, const GLfloat *v ) {
}

MultiTexCoord3i( GLenum target, GLint s, GLint t, GLint r ) {
}

MultiTexCoord3iv( GLenum target, const GLint *v ) {
}

MultiTexCoord3s( GLenum target, GLshort s, GLshort t, GLshort r ) {
}

MultiTexCoord3sv( GLenum target, const GLshort *v ) {
}

MultiTexCoord4d( GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q ) {
}

MultiTexCoord4dv( GLenum target, const GLdouble *v ) {
}

MultiTexCoord4f( GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q ) {
}

MultiTexCoord4fv( GLenum target, const GLfloat *v ) {
}

MultiTexCoord4i( GLenum target, GLint s, GLint t, GLint r, GLint q ) {
}

MultiTexCoord4iv( GLenum target, const GLint *v ) {
}

MultiTexCoord4s( GLenum target, GLshort s, GLshort t, GLshort r, GLshort q ) {
}

MultiTexCoord4sv( GLenum target, const GLshort *v ) {
}


LoadTransposeMatrixd( const GLdouble m[16] ) {
}

LoadTransposeMatrixf( const GLfloat m[16] ) {
}

MultTransposeMatrixd( const GLdouble m[16] ) {
}

MultTransposeMatrixf( const GLfloat m[16] ) {
}

SampleCoverage( GLclampf value, GLboolean invert ) {
}

ActiveTextureARB(GLenum texture) {
}
ClientActiveTextureARB(GLenum texture) {
}
MultiTexCoord1dARB(GLenum target, GLdouble s) {
}
MultiTexCoord1dvARB(GLenum target, const GLdouble *v) {
}
MultiTexCoord1fARB(GLenum target, GLfloat s) {
}
MultiTexCoord1fvARB(GLenum target, const GLfloat *v) {
}
MultiTexCoord1iARB(GLenum target, GLint s) {
}
MultiTexCoord1ivARB(GLenum target, const GLint *v) {
}
MultiTexCoord1sARB(GLenum target, GLshort s) {
}
MultiTexCoord1svARB(GLenum target, const GLshort *v) {
}
MultiTexCoord2dARB(GLenum target, GLdouble s, GLdouble t) {
}
MultiTexCoord2dvARB(GLenum target, const GLdouble *v) {
}
MultiTexCoord2fARB(GLenum target, GLfloat s, GLfloat t) {
}
MultiTexCoord2fvARB(GLenum target, const GLfloat *v) {
}
MultiTexCoord2iARB(GLenum target, GLint s, GLint t) {
}
MultiTexCoord2ivARB(GLenum target, const GLint *v) {
}
MultiTexCoord2sARB(GLenum target, GLshort s, GLshort t) {
}
MultiTexCoord2svARB(GLenum target, const GLshort *v) {
}
MultiTexCoord3dARB(GLenum target, GLdouble s, GLdouble t, GLdouble r) {
}
MultiTexCoord3dvARB(GLenum target, const GLdouble *v) {
}
MultiTexCoord3fARB(GLenum target, GLfloat s, GLfloat t, GLfloat r) {
}
MultiTexCoord3fvARB(GLenum target, const GLfloat *v) {
}
MultiTexCoord3iARB(GLenum target, GLint s, GLint t, GLint r) {
}
MultiTexCoord3ivARB(GLenum target, const GLint *v) {
}
MultiTexCoord3sARB(GLenum target, GLshort s, GLshort t, GLshort r) {
}
MultiTexCoord3svARB(GLenum target, const GLshort *v) {
}
MultiTexCoord4dARB(GLenum target, GLdouble s, GLdouble t, GLdouble r, GLdouble q) {
}
MultiTexCoord4dvARB(GLenum target, const GLdouble *v) {
}
MultiTexCoord4fARB(GLenum target, GLfloat s, GLfloat t, GLfloat r, GLfloat q) {
}
MultiTexCoord4fvARB(GLenum target, const GLfloat *v) {
}
MultiTexCoord4iARB(GLenum target, GLint s, GLint t, GLint r, GLint q) {
}
MultiTexCoord4ivARB(GLenum target, const GLint *v) {
}
MultiTexCoord4sARB(GLenum target, GLshort s, GLshort t, GLshort r, GLshort q) {
}
MultiTexCoord4svARB(GLenum target, const GLshort *v) {
}
*/
