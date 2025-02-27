import router from "@/routes"
import { NamesRouter } from "@/routes/routesNames"

export const navigateToDetails = (id:string) =>{
  router.push({
    name:NamesRouter.MailDetail,
    params:{id}
  })
}
