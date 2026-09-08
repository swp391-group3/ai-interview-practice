"""PASS 1 reviewed primitive structure. Run with Blender --background --factory-startup --python this_file.py.
Writes a separate pass1-reproduced.blend; never overwrites the reviewed working scene.
Consolidated only after staged viewport, A/B/C render and A crop review.
"""
import bpy, math
from mathutils import Vector
from pathlib import Path
ROOT=str(Path(__file__).resolve().parents[1])
if any(o.name not in {'Cube','Camera','Light'} for o in bpy.context.scene.objects):
    raise RuntimeError('Run in a fresh Blender default scene; preserve reviewed work.')
bpy.ops.object.select_all(action='SELECT')
bpy.ops.object.delete(use_global=False)
for c in list(bpy.data.collections):
    if c.name=='Collection': bpy.data.collections.remove(c)
COLS=['00_GUIDES','10_ARCHITECTURE','20_FURNITURE','30_PROPS_BLOCKOUT','40_CANDIDATE_PROXY','50_SCREEN','60_LIGHTS','70_CAMERAS']
for n in COLS:
    if not bpy.data.collections.get(n): bpy.context.scene.collection.children.link(bpy.data.collections.new(n))
def mat(n,c):
    m=bpy.data.materials.new(n); m.diffuse_color=(*c,1); m.use_nodes=True
    p=m.node_tree.nodes.get('Principled BSDF'); p.inputs['Base Color'].default_value=(*c,1); p.inputs['Roughness'].default_value=.78
    return m
wall=mat('Clay | chalk',(.73,.70,.64))
wood=mat('Clay | warm wood',(.32,.22,.145))
lightwood=mat('Clay | storage',(.52,.43,.32))
dark=mat('Clay | graphite',(.055,.065,.068))
cloth=mat('Clay | candidate',(.12,.19,.19))
skin=mat('Clay | proxy head hands',(.36,.32,.28))
paper=mat('Clay | paper',(.78,.76,.69))
accent=mat('Clay | terracotta',(.35,.16,.10))
def link(o,n,col,m):
    o.name=n
    for c in list(o.users_collection): c.objects.unlink(o)
    bpy.data.collections[col].objects.link(o)
    if m: o.data.materials.append(m)
    return o
def box(n,loc,dim,m=wall,col='10_ARCHITECTURE',bevel=.015):
    bpy.ops.mesh.primitive_cube_add(size=1,location=loc); o=link(bpy.context.object,n,col,m); o.dimensions=dim
    bpy.ops.object.transform_apply(location=False,rotation=False,scale=True)
    if bevel:
        mod=o.modifiers.new('Blockout edge','BEVEL'); mod.width=bevel; mod.segments=2
        o.modifiers.new('Weighted normals','WEIGHTED_NORMAL')
    return o
def ell(n,loc,scale,m=cloth,col='40_CANDIDATE_PROXY'):
    bpy.ops.mesh.primitive_uv_sphere_add(segments=24,ring_count=12,location=loc); o=link(bpy.context.object,n,col,m); o.scale=scale
    for p in o.data.polygons:p.use_smooth=True
    return o
def rod(n,a,b,r,m=dark,col='20_FURNITURE'):
    a,b=Vector(a),Vector(b)
    bpy.ops.mesh.primitive_cylinder_add(vertices=20,radius=r,depth=(b-a).length,location=(a+b)/2)
    o=link(bpy.context.object,n,col,m); o.rotation_euler=(b-a).to_track_quat('Z','Y').to_euler(); return o
def aim(o,p):o.rotation_euler=(Vector(p)-o.location).to_track_quat('-Z','Y').to_euler()
def cam(n,loc,target,lens):
    d=bpy.data.cameras.new(n); o=bpy.data.objects.new(n,d); bpy.data.collections['70_CAMERAS'].objects.link(o);o.location=loc;d.lens=lens;aim(o,target);return o
scene=bpy.context.scene
scene.unit_settings.system='METRIC'
scene.render.engine='BLENDER_EEVEE'
scene.render.resolution_x=1600;scene.render.resolution_y=900;scene.render.resolution_percentage=100
scene.world.color=(.18,.18,.18)
box('Floor | 4.2 x 4.0 m',(0,0,-.09),(4.2,4,.18),lightwood)
box('Back wall',(0,1.95,1.4),(4.2,.14,2.8))
box('Left window sill wall',(-2.03,0,.43),(.14,4,.86))
box('Left window header',(-2.03,0,2.62),(.14,4,.36))
box('Left window rear pier',(-2.03,1.58,1.65),(.14,.7,1.6))
box('Left window front pier',(-2.03,-1.72,1.65),(.14,.56,1.6))
box('Window sill',(-1.97,-.05,.88),(.27,2.9,.06),paper)
box('Window mullion',(-2.02,-.12,1.66),(.08,.055,1.55),paper)
box('Blind folded header',(-1.92,-.1,2.39),(.10,2.85,.16),paper)
box('Credenza',(0.85,1.63,.36),(1.9,.46,.72),lightwood,'20_FURNITURE')
for x in [.23,.87,1.49]:box('Storage door',(x,1.389,.36),(.60,.02,.63),wall,'20_FURNITURE')
a=cam('Camera A | three-quarter',(-2.35,-3.55,1.65),(.05,.4,1.02),50)
scene.camera=a
for area in (bpy.context.screen.areas if bpy.context.screen else []):
    if area.type=='VIEW_3D':
        area.spaces.active.region_3d.view_perspective='CAMERA'
        area.spaces.active.overlay.show_overlays=False
        area.spaces.active.shading.color_type='MATERIAL'


box('Right wall',(2.03,0,1.4),(.14,4,2.8))
a.location=(-1.65,-3.45,1.65);aim(a,(.05,.45,1.08))


F='20_FURNITURE';P='30_PROPS_BLOCKOUT'
box('Desk | 160 x 78 cm',(.30,.35,.745),(1.60,.78,.065),wood,F)
for x in [-.38,.98]:
    for y in [.05,.66]:box('Desk leg',(x,y,.36),(.055,.055,.72),dark,F)
box('Chair seat',(.66,-.48,.465),(.49,.48,.10),dark,F,.045)
box('Chair low back',(.66,-.70,.73),(.48,.075,.43),dark,F,.04)
rod('Chair pedestal',(.66,-.48,.12),(.66,-.48,.42),.045)
for dx,dy in [(-.25,-.22),(.25,-.22),(-.26,.2),(.26,.2)]:
    rod('Chair base',(.66,-.48,.13),(.66+dx,-.48+dy,.07),.022)
box('Monitor foot',(.19,.54,.796),(.27,.18,.025),dark,F)
rod('Monitor stem',(.19,.60,.80),(.19,.60,1.05),.027)
box('Monitor bezel',(.19,.60,1.20),(.66,.047,.40),dark,F,.014)
box('Screen | replaceable surface',(.19,.573,1.20),(.615,.008,.350),paper,'50_SCREEN',.004)
box('Keyboard',(.31,.11,.795),(.40,.145,.025),dark,P,.008)
box('Open notebook left',(-.20,.015,.791),(.16,.24,.013),paper,P,.003)
box('Open notebook right',(-.035,.015,.793),(.16,.24,.013),paper,P,.003)
o=box('Resume | blank placeholder',(-.34,.31,.788),(.21,.297,.003),paper,P,.001);o.rotation_euler.z=-.12
rod('Pen',(-.36,-.08,.80),(-.36,.07,.80),.004,dark,P)
rod('Water vessel',(.84,.44,.78),(.84,.44,.92),.038,paper,P)
box('Reference book',(-.31,.59,.802),(.23,.16,.035),accent,P,.004)
for x in [.74,.92]:ell('Headphone ear cup',(x,-.005,.815),(.035,.05,.025),dark,P)
rod('Headphone bridge',(.74,.035,.82),(.92,.035,.82),.012,dark,P)
box('Background book mass',(1.47,1.60,.81),(.30,.22,.15),accent,P)
rod('Practical lamp stem',(.93,1.62,.72),(.93,1.62,1.18),.018,dark,P)
bpy.ops.mesh.primitive_cone_add(vertices=32,radius1=.15,radius2=.10,depth=.20,location=(.93,1.62,1.20));link(bpy.context.object,'Practical shade',P,paper)


C='40_CANDIDATE_PROXY'
ell('Pelvis',(.66,-.43,.59),(.18,.16,.13))
o=ell('Torso attentive lean',(.64,-.37,.93),(.225,.135,.32));o.rotation_euler.x=-.10
ell('Neck',(.62,-.31,1.235),(.062,.064,.09),skin)
o=ell('Head | screen gaze',(.60,-.275,1.385),(.10,.115,.145),skin);o.rotation_euler.z=.32
def limb(n,a,b,r,m=cloth):
    o=ell(n,(Vector(a)+Vector(b))/2,(r,r,(Vector(b)-Vector(a)).length/2+r*.35),m)
    o.rotation_euler=(Vector(b)-Vector(a)).to_track_quat('Z','Y').to_euler()
limb('Left upper arm',(.43,-.35,1.12),(.39,-.31,.87),.067)
limb('Left forearm',(.39,-.31,.87),(.32,.02,.815),.051)
ell('Left hand | keyboard',(.31,.065,.816),(.045,.075,.023),skin)
limb('Right upper arm',(.85,-.34,1.11),(.88,-.43,.87),.069)
limb('Right forearm',(.88,-.43,.87),(.84,-.10,.805),.050)
ell('Right hand | resting',(.84,-.055,.809),(.044,.067,.024),skin)
limb('Left thigh',(.53,-.41,.58),(.49,.015,.53),.086)
limb('Right thigh',(.78,-.42,.58),(.83,-.035,.53),.089)
limb('Left shin',(.49,.015,.53),(.46,.07,.13),.063)
limb('Right shin',(.83,-.035,.53),(.86,-.12,.13),.063)
ell('Left foot',(.46,.16,.08),(.07,.14,.06),dark)
ell('Right foot',(.86,-.025,.08),(.07,.14,.06),dark)
aim(a,(.30,.40,1.06))


S='50_SCREEN'
ui=bpy.data.materials.new('Clay | screen muted blue grey');ui.diffuse_color=(.25,.34,.36,1);ui.use_nodes=True
p=ui.node_tree.nodes.get('Principled BSDF');p.inputs['Base Color'].default_value=(.25,.34,.36,1);p.inputs['Roughness'].default_value=.9
bpy.data.objects['Screen | replaceable surface'].data.materials.clear();bpy.data.objects['Screen | replaceable surface'].data.materials.append(ui)
box('Placeholder | interviewer panel',(.035,.564,1.205),(.255,.006,.285),lightwood,S,.005)
ell('Placeholder | interviewer head',(.035,.555,1.25),(.036,.005,.042),paper,S)
ell('Placeholder | interviewer shoulders',(.035,.555,1.166),(.075,.005,.042),paper,S)
box('Placeholder | interaction area',(.321,.564,1.215),(.245,.006,.235),paper,S,.004)
box('Placeholder | neutral interaction bar',(.315,.557,1.112),(.175,.004,.012),ui,S,.002)
box('Placeholder | control region',(.19,.557,1.045),(.17,.004,.012),dark,S,.002)


def area(n,loc,target,power,color,size):
    d=bpy.data.lights.new(n,'AREA');d.energy=power;d.color=color;d.shape='DISK';d.size=size
    o=bpy.data.objects.new(n,d);bpy.data.collections['60_LIGHTS'].objects.link(o);o.location=loc;aim(o,target);return o
area('Daylight | broad left window',(-1.88,-.15,2.05),(.4,.35,.7),230,(.86,.93,1),2.4)
area('Interior | gentle fill',(.2,-1.5,2.4),(.5,.2,.8),65,(1,.92,.81),2.6)
area('Practical | warm pool',(.93,1.61,1.08),(.9,1.55,.72),12,(1,.65,.35),.25)
scene.world.use_nodes=True
scene.world.node_tree.nodes['Background'].inputs['Color'].default_value=(.63,.68,.73,1)
scene.world.node_tree.nodes['Background'].inputs['Strength'].default_value=.22
scene.view_settings.view_transform='AgX'
scene.render.image_settings.file_format='PNG'
scene.render.filepath=ROOT+'/renders/blockout/camera-a-pass1.png'
scene.render.film_transparent=False




for n in ['Practical lamp stem','Practical shade','Practical | warm pool']:
    bpy.data.objects[n].location.x+=.39
P='30_PROPS_BLOCKOUT'
for x,y in [(-.23,.06),(-.23,-.015),(-.065,.04)]:
    for dx,dy,sx,sy in [(0,-.016,.045,.002),(0,.016,.045,.002),(-.022,0,.002,.032),(.022,0,.002,.032)]:
        box('Notebook | diagram box',(x+dx,y+dy,.802),(sx,sy,.0015),dark,P,0)
box('Notebook | connection',(-.146,.04,.803),(.12,.002,.0015),dark,P,0)
for i in range(4):
    box('Resume | abstract rule',(-.34,.37-i*.032,.791),(.12 if i else .08,.003,.001),lightwood,P,0)
b=cam('Camera B | direct shoulder',(-.58,-2.65,1.60),(.29,.40,1.12),50)
c=cam('Camera C | environmental',(-2.00,-4.60,1.85),(.04,.45,1.02),45)
scene.camera=a

scene.render.filepath=ROOT+'/renders/blockout/camera-a-pass1.png'



g=bpy.data.objects.new('Crop guide | A 4x5 x32-77 percent',None)
bpy.data.collections['00_GUIDES'].objects.link(g);g.hide_render=True;g.hide_viewport=True
g['desktop_aspect']='16:9';g['laptop_crop_x']=[.078125,.921875];g['portrait_crop_x']=[.32,.77];g['portrait_note']='Keep head, keyboard hand, full screen and open notebook; chair and resume may crop.'
scene.camera=a
scene.render.use_border=False;scene.render.use_crop_to_border=False
scene.render.filepath=ROOT+'/renders/blockout/camera-a-pass1.png'



bpy.ops.wm.save_as_mainfile(filepath=ROOT+'/scenes/pass1-reproduced.blend')
