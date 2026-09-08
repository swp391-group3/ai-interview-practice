# original vs clone · 克隆评估报告

## 结论
- 原站 URL: https://open-design.ai/
- 克隆 URL: http://127.0.0.1:8123/reference-study.html
- 自动推断复杂度: L5
- 复刻模式建议: 技术拆解 / 忠实复刻优先
- 自动报告边界: 结构、数量、框架、console 可自动比；传入 visual-diff 后可纳入像素差异分。内容残留和法务仍需审计。

## 技术信号
| 项目 | 原站 | 克隆站 |
|---|---|---|
| title | OpenDesign — Best Open Source Claude Design Alternative | Reference Study — Design & Motion Grammar |
| lang | en | en |
| frameworks | none | none |
| scrollHeight | 13108 | 13193 |
| h1 | The Vibe Design Workspace for your brand.One design system. Every prototype, slide, marketing image,video, and dashboard stays on-brand. | A design system that makes every screen feel considered. |

## 数量对比
| 指标 | 原站 | 克隆站 | 自动评分 |
|---|---:|---:|---:|
| sections | 13 | 30 | 2/5 |
| links | 178 | 38 | 1/5 |
| images | 60 | 0 | 1/5 |
| video | 1 | 0 | 1/5 |
| canvas | 2 | 0 | 1/5 |
| forms | 1 | 1 | 5/5 |
| buttons | 8 | 17 | 2/5 |
| inputs | 4 | 1 | 1/5 |
| interactive | 207 | 56 | 1/5 |
| scripts | 8 | 0 | 1/5 |

## 复刻评分
- 源证据: 3/5
- 结构保真: 1/5
- 视觉保真: 2/5
- 动效/交互: 1/5
- 响应式: 4/5
- 功能完整: 2/5
- 内容替换: 需人工看文案残留
- 法务/部署风险: 需人工核查 license / 素材

## Console
- 原站 console errors: 0
- 克隆 console errors: 0
- 原站 page errors: 0
- 克隆 page errors: 0

## 路由覆盖
- 原站路由: 25
- 克隆路由: 4
- 覆盖率: 0%
- 原站 route map: RECON/routes/original-route-map.json
- 克隆 route map: RECON/routes-clone/clone-route-map.json
- 缺失路由: /, /pricing, /codex-plugin, /solutions, /solutions/prototype, /solutions/dashboard, /solutions/slides, /solutions/image, /solutions/video, /solutions/design-system, /solutions/solo-builder, /solutions/designer, /solutions/engineering, /solutions/product-managers, /solutions/marketing, /solutions/ai-wireframe-generator, /solutions/ai-ui-generator, /solutions/ai-prototype-generator, /solutions/ai-landing-page-generator, /solutions/design-to-code, /solutions/figma-to-code, /solutions/screenshot-to-code, /solutions/html-to-ppt, /agents, /agents/deepseek-harness-design
- 额外路由: /reference-study.html, /REFERENCE-AUDIT.md, /MOTION-AUDIT.md, /TOKENS.css


## 交互覆盖
- 原站可见交互目标: 80
- 克隆可见交互目标: 45
- 原站 canvas 目标: 1
- 克隆 canvas 目标: 0
- 原站 changed actions: 8/23
- 克隆 changed actions: 15/22
- 原站 interaction probe: RECON/interactions/original-interactions.json
- 克隆 interaction probe: RECON/interactions-clone/clone-interactions.json
- 判断: 交互数量信号不一致，需要检查缺失状态或过度实现。


## 截图证据
- 原站侦察: RECON/original-recon.json
- 克隆侦察: RECON/clone-recon.json
- 像素差异: RECON/visual-diff-1440.json
- 像素差异率: 0.23629128873055577
- 原站截图: screenshots/original-1440.png, screenshots/original-768.png, screenshots/original-390.png
- 克隆截图: screenshots/clone-1440.png, screenshots/clone-768.png, screenshots/clone-390.png

## 已知缺口
- 未传入 visual-diff 时，视觉保真需要打开截图人工确认。
- 法务、素材授权、品牌替换完整度需要人工核查。
