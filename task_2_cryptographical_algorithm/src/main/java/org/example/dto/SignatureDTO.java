package org.example.dto;

import lombok.*;
import org.example.sigalgorithms.Signature;

import java.io.Serializable;
import java.math.BigInteger;
import java.util.List;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@ToString
public class SignatureDTO implements Serializable {
    private ECPointDTO keyImage;
    private List<BigInteger> cList;
    private List<BigInteger> rList;

    public SignatureDTO(Signature signature){
        this.keyImage = new ECPointDTO(signature.getKeyImage().getXCoord().toBigInteger(),
                signature.getKeyImage().getYCoord().toBigInteger());
        this.cList = signature.getCList();
        this.rList = signature.getRList();
    }
}
