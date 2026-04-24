<template>
  <div class="game-page">
    <web-navbar :no-scroll="true"/>
    <div class="container">
      <br>
      <div class="game-sections">
        <div v-for="section in gameSections" :key="section.type" class="game-section">
          <h2 class="section-title">{{ section.title }}</h2>
          <div class="game-grid">
            <div 
              v-for="game in section.games" 
              :key="game.id" 
              class="game-card"
            >
              <div class="game-card-inner">
                <img :src="game.image" :alt="game.title" class="game-image">
                <div class="game-overlay">
                  <div class="game-info">
                    <h3 class="game-title">{{ game.title }}</h3>
                    <p class="game-description">{{ game.description }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import WebNavbar from '@/components/layout/WebNavbar.vue'

interface Game {
  id: number
  title: string
  image: string
  description: string
}

interface GameSection {
  type: string
  title: string
  games: Game[]
}


const gameSections = ref<GameSection[]>([
  {
    type: 'action',
    title: 'ARPG类',
    games: [
      { id: 1, title: '永劫无间', image: 'https://tse4-mm.cn.bing.net/th/id/OIP-C.Sm7_MosqgzR9HebrKfXoBwHaEK?w=277&h=180&c=7&r=0&o=7&dpr=1.3&pid=1.7&rm=3', description: '一款以中国风为背景的武侠动作竞技游戏，融合了武侠元素与大逃杀玩法，玩家可以使用各种武器和技能进行战斗。' },
      { id: 2, title: '怪物猎人世界', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/img/acbd7ecb60e4612af31ba23e1f853013~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036452&x-signature=DNwpsiyWLDLq0YTuSzokTYzfDos%3D', description: '一款开放世界狩猎游戏，玩家需要探索广阔的地图，狩猎各种巨大的怪物，收集素材制作装备，不断提升自己的狩猎技能。' },
      { id: 3, title: '鬼泣', image: 'https://p3-flow-imagex-sign.byteimg.com/labis/image/1ede49aad40c508e428861a82d51554f~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=2026042017203190D2C22BC74CB04F6AB1&rrcfp=cee388b0&x-expires=2092036839&x-signature=RN41MubKobYsWdG7QMbOT7gDr1o%3D', description: '一款经典的动作游戏系列，以华丽的战斗系统和流畅的连击著称，玩家扮演恶魔猎人但丁，与各种恶魔进行激烈的战斗。' }
    ]
  },
  {
    type: 'rpg',
    title: '二游类',
    games: [
      { id: 4, title: '原神', image: 'https://p3-flow-imagex-sign.byteimg.com/labis/383c883f3c91cb2528f894b20a4422ba~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036452&x-signature=UjKCL48mT53FhrFkR%2FO7GUnMijY%3D', description: '一款开放世界冒险游戏，主角开局失忆，玩家可以探索提瓦特大陆，邂逅各种角色，使用元素之力进行战斗，体验丰富的剧情和任务。' },
      { id: 5, title: '鸣潮', image: 'https://p3-flow-imagex-sign.byteimg.com/labis/image/d1a678016f1def56027ff5e323609634~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036452&x-signature=3B%2BAazTi4uEl91pmpL4%2FjfE4RGc%3D', description: '一款科幻题材的开放世界动作 RPG，主角开局失忆，玩家可以在充满未来感的世界中自由探索，使用各种武器和技能与敌人战斗。' },
      { id: 6, title: '异环', image: 'https://ts2.tc.mm.bing.net/th/id/OIP-C.HS6oYYdb1TojcpQ6aNMy1AHaHa?w=108&h=108&c=1&bgcl=f0990a&r=0&o=7&dpr=1.3&pid=ImgRC&rm=3', description: '一款以现代都市风格结合各种奇怪异物为背景的角色扮演游戏，主角开局失忆，玩家需要和异象调查局的伙伴一起维护城市治安，解开神秘的谜团。' }
    ]
  },
  {
    type: 'strategy',
    title: '策略类',
    games: [
      { id: 7, title: '三国杀', image: '/image/三国杀.jpg', description: '一款以三国为历史题材的策略卡牌游戏，常年霸占steam差评榜榜首，虽然因全服共享保底，新手教程教你怎么充值等原因被不少玩家诟病，但不可否认他是我们学生时代最火的线下卡牌游戏之一。' },
      { id: 8, title: '皇室战争', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/img/56b03cf180e0fb596633493b54a18315~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036453&x-signature=z06%2Ffrgeo36VRFCK3ecYDByI6gY%3D', description: '一款快节奏的策略卡牌游戏，玩家需要收集和升级各种卡牌，组建强大的卡组，与其他玩家进行实时对战。' },
    ]
  },
  {
    type: 'three a',
    title: '3A类',
    games: [
      { id: 9, title: '荒野大镖客2', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/image/1e196c247af50b3ba6ac199a0df9a443~tplv-a9rns2rl98-pc_smart_face_crop-v1:360:270.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036453&x-signature=X5WekYXPg0anq%2FnEqM7y7jsQxpI%3D', description: '一款开放世界动作冒险游戏，以19世纪末的美国西部为背景，玩家扮演亚瑟·摩根，体验西部牛仔的生活，参与各种任务和活动。' },
      { id: 10, title: 'gta5', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/img/1448e65f3cfe352641a7918f49e9667b~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036453&x-signature=ClW3%2BrboDN5ajW8dBdGeSXeTALs%3D', description: '一款开放世界动作冒险游戏，玩家可以在洛圣都自由探索，参与各种任务，驾驶各种车辆，体验丰富的游戏内容。' },
      
    ]
  },
  {
    type: 'moba',
    title: 'Moba类',
    games: [
      { id: 11, title: '王者荣耀', image: 'https://tse3-mm.cn.bing.net/th/id/OIP-C.F7DJPtFqHgEyg8IFhzsAOQHaEK?w=318&h=180&c=7&r=0&o=7&dpr=1.3&pid=1.7&rm=3', description: '一款多人在线战术竞技游戏，玩家通过控制英雄，与队友合作，摧毁敌方基地，是目前国内最受欢迎的MOBA游戏之一。' },
      { id: 12, title: '英雄联盟', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/image/1aa99bf78a2013d2310aa9c82c5fb5d2~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420172227308216BEF92AAD550E5C&rrcfp=cee388b0&x-expires=2092036954&x-signature=k7C6S40pwwqDzvg1CF0hAtl8GrI%3D', description: '一款全球流行的MOBA游戏，玩家控制具有不同技能的英雄，与队友合作对抗敌方团队，目标是摧毁对方的基地。' },
      { id: 13, title: '刀塔2', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/image/99788bab080426d3da00092eec4df8e0~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420172227308216BEF92AAD550E5C&rrcfp=cee388b0&x-expires=2092036954&x-signature=exwjHWP6BA2SQVal3ydeCpYmWVM%3D', description: '一款基于《魔兽争霸3》地图编辑器制作的MOBA游戏，拥有丰富的英雄池和复杂的游戏机制，是电子竞技的重要项目。' }
    ]
  },
  {
    type: 'fps',
    title: 'fps类',
    games: [
      { id: 14, title: '求生之路2', image: 'https://p3-flow-imagex-sign.byteimg.com/labis/30348229164f12cdc9106cbfe42a782b~tplv-a9rns2rl98-pc_smart_face_crop-v1:512:384.image?lk3s=8e244e95&rcl=20260420171351C58FAFA34722BC4BB7CD&rrcfp=cee388b0&x-expires=2092036453&x-signature=vLCFso1yXwKilWgqBNH6zOVAtAQ%3D', description: '一款合作生存射击游戏，玩家需要与队友合作，在充满僵尸的世界中生存，通过各种武器和道具对抗来袭的僵尸潮。' },
      { id: 15, title: 'Counter-Strike: Global Offensive', image: 'https://p3-flow-imagex-sign.byteimg.com/isp-i18n-media/img/02697b21df635dd72c32ccede31a770f~tplv-a9rns2rl98-pc_smart_face_crop-v1:286:215.image?lk3s=8e244e95&rcl=202604201721245C14BB1DE759752CE176&rrcfp=cee388b0&x-expires=2092036889&x-signature=xB4p0brMHGn0PpY6QPcytSgVH0o%3D', description: '一款经典的第一人称射击游戏，玩家分为反恐精英和恐怖分子两个阵营，通过团队合作和个人技巧完成各种任务。' }
    ]
  },
  {
    type:'galgame',
    title:'视觉小说类',
    games:[    
      {id:16,title:'魔法少女的魔女审判',image:'https://tse3-mm.cn.bing.net/th/id/OIP-C.CAMrAxP0wZkIPPOPofiK-gHaDt?w=344&h=175&c=7&r=0&o=7&dpr=1.3&pid=1.7&rm=3',description:'一款魔法少女题材的视觉小说游戏，玩家将体验魔法少女与魔女之间的战斗，通过选择不同的选项影响故事的发展。'}
    ]
  },
  {
    type:'mmo',
    title:'MMO类',
    games:[
      {id:17,title:'逆水寒',image:'/image/逆水寒.jpg',description:'网易出品的国风武侠 MMORPG，北宋背景，画质精美，开放世界，主打休闲社交与江湖剧情。'},
      {id:18,title:'燕云十六声',image:'/image/燕云十六声.jpg',description:'高自由度开放世界武侠游戏，五代十国背景，无门派无等级，侧重自由武学与动作战斗。'}
  ]}
])
</script>

<style scoped lang="scss">
.game-page {
  padding: 20px;
  min-height: 100vh;
  background-color: #FFF8E7;
  .container {
    max-width: 1400px;
    margin: 0 auto;

    .page-title {
      font-size: 36px;
      color: #333;
      text-align: center;
      margin-bottom: 40px;
      font-weight: bold;
    }

    .game-sections {
      .game-section {
        margin-bottom: 60px;

        .section-title {
          font-size: 28px;
          color: #333;
          margin-bottom: 30px;
          padding-left: 20px;
          border-left: 4px solid #ffd700;
          font-weight: bold;
        }

        .game-grid {
          display: grid;
          grid-template-columns: repeat(3, 1fr);
          gap: 25px;

          .game-card {
            position: relative;
            aspect-ratio: 1 / 1;
            border-radius: 12px;
            overflow: hidden;
            box-shadow: 0 4px 15px rgba(0, 0, 0, 0.3);
            transition: all 0.3s ease;
            cursor: pointer;

            &:hover {
              transform: scale(1.05);
              box-shadow: 0 8px 25px rgba(0, 0, 0, 0.4);
              z-index: 10;

              .game-card-inner {
                .game-image {
                  filter: blur(3px);
                  transform: scale(1.1);
                }

                .game-overlay {
                  opacity: 1;
                }
              }
            }

            .game-card-inner {
              position: relative;
              width: 100%;
              height: 100%;

              .game-image {
                width: 100%;
                height: 100%;
                object-fit: cover;
                transition: all 0.3s ease;
              }

              .game-overlay {
                position: absolute;
                top: 0;
                left: 0;
                width: 100%;
                height: 100%;
                background: linear-gradient(to top, rgba(0, 0, 0, 0.9), rgba(0, 0, 0, 0.6), transparent);
                display: flex;
                align-items: flex-end;
                justify-content: center;
                padding: 20px;
                opacity: 0;
                transition: opacity 0.3s ease;

                .game-info {
                  text-align: center;
                  color: white;

                  .game-title {
                    font-size: 20px;
                    margin-bottom: 10px;
                    font-weight: bold;
                    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
                  }

                  .game-description {
                    font-size: 14px;
                    line-height: 1.6;
                    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.5);
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}

@media (max-width: 768px) {
  .game-page {
    .container {
      .page-title {
        font-size: 28px;
        margin-bottom: 30px;
      }

      .game-sections {
        .game-section {
          margin-bottom: 40px;

          .section-title {
            font-size: 22px;
            margin-bottom: 20px;
          }

          .game-grid {
            grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
            gap: 15px;

            .game-card {
              &:hover {
                .game-card-inner {
                  .game-image {
                    filter: blur(2px);
                  }
                }
              }

              .game-card-inner {
                .game-overlay {
                  padding: 10px;

                  .game-info {
                    .game-title {
                      font-size: 16px;
                      margin-bottom: 5px;
                    }

                    .game-description {
                      font-size: 12px;
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
}
</style>