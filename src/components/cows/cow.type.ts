export type Cow = {
    id: string
    birthdate: Date
    colour: string
    gender: string
    breed: string
    motherId: string
    farmerId: string
    fatherBreed: string
    isPregnant: boolean
    ovulation: Date
    pregnancies: Pregnancy[]
    inseminations: Insemination[]
}

export type Insemination = {
    date: Date
    breed: string
    IsArtificial: boolean
}
export type Pregnancy = {
    detectedAt: Date
    firstDay: Date
    lastDay: Date
}
