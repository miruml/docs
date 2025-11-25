export const FramedImage = ({
    image,
    background,
    alt = "Framed content",
    width = "100%",
    borderWidth = "18px",
    outerRadius = "12px",
    innerRadius = "10px",
}) => {
    return (
        <div
            style={{
                padding: `${borderWidth}`,
                backgroundImage: background ? `url(${background})` : undefined,
                backgroundSize: "cover",
                backgroundPosition: "center",
                borderRadius: outerRadius,
                width: width,
            }}
        >
            <div className="overflow-hidden" >
                <img
                    src={image}
                    alt={alt}
                    style={{
                        borderRadius: innerRadius,
                        margin: `0 0`,
                        width: "100%",
                        display: "block",
                    }}
                />
            </div>
        </div>
    );
};
